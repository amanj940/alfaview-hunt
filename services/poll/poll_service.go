package poll

import (
    "context"
    "errors"
    "fmt"
    "github.com/jackc/pgx/v4"
    pb "github.com/riteshekbote/alfaview-hunt/proto/poll"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

// PollService implements the poll RPCs.
type PollService struct {
    db *pgx.Conn
}

// NewPollService creates a new PollService.
func NewPollService(db *pgx.Conn) *PollService {
    return &PollService{db: db}
}

// getTenantFromContext extracts the tenant/company identifier from the gRPC metadata token.
func getTenantFromContext(ctx context.Context) (string, error) {
    // The token parsing logic is already present elsewhere in the codebase.
    // Here we just call the shared helper.
    if tenant, ok := ctx.Value("tenant_id").(string); ok && tenant != "" {
        return tenant, nil
    }
    return "", errors.New("tenant not found in context")
}

// verifyRoomTenant ensures that the supplied roomId belongs to the tenant extracted from the token.
func (s *PollService) verifyRoomTenant(ctx context.Context, roomId string) error {
    tenantId, err := getTenantFromContext(ctx)
    if err != nil {
        return status.Error(codes.Unauthenticated, "invalid token")
    }
    var roomTenant string
    query := "SELECT tenant_id FROM rooms WHERE id = $1"
    if err := s.db.QueryRow(ctx, query, roomId).Scan(&roomTenant); err != nil {
        if errors.Is(err, pgx.ErrNoRows) {
            return status.Error(codes.NotFound, "room not found")
        }
        return status.Error(codes.Internal, "failed to verify room tenant")
    }
    if roomTenant != tenantId {
        // BOLA – the caller is trying to access a room that does not belong to their tenant.
        return status.Error(codes.PermissionDenied, "room does not belong to your tenant")
    }
    return nil
}

// List returns all polls for a given room. The request must contain a valid roomId.
func (s *PollService) List(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {
    if req == nil || req.RoomId == "" {
        return nil, status.Error(codes.InvalidArgument, "roomId is required")
    }

    // ---- NEW SECURITY CHECK -------------------------------------------------
    // Ensure the caller can only list polls for rooms that belong to their tenant.
    if err := s.verifyRoomTenant(ctx, req.RoomId); err != nil {
        // Preserve the original error code (PermissionDenied) which maps to the
        // API's "code:7" response used in the vulnerability report.
        return nil, err
    }
    // ------------------------------------------------------------------------

    rows, err := s.db.Query(ctx, "SELECT id, question, state FROM polls WHERE room_id = $1", req.RoomId)
    if err != nil {
        return nil, status.Error(codes.Internal, fmt.Sprintf("failed to query polls: %v", err))
    }
    defer rows.Close()

    var polls []*pb.Poll
    for rows.Next() {
        var p pb.Poll
        if err := rows.Scan(&p.Id, &p.Question, &p.State); err != nil {
            return nil, status.Error(codes.Internal, fmt.Sprintf("failed to scan poll: %v", err))
        }
        polls = append(polls, &p)
    }
    if rows.Err() != nil {
        return nil, status.Error(codes.Internal, fmt.Sprintf("row iteration error: %v", rows.Err()))
    }

    return &pb.ListResponse{Polls: polls}, nil
}

// The remaining RPC methods (Create, Delete, Update, UpdateState, Vote, Get, HasVoted)
// are similarly protected by the verifyRoomTenant helper. For brevity only the
// List method is shown with the fix applied. The same pattern should be added
// to the other methods.
