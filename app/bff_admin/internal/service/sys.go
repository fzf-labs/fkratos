package service

import (
	"context"
	"fkratos/api/paginator"
	sysV1 "fkratos/api/rpc_sys/v1"

	pb "fkratos/api/bff_admin/v1"

	"github.com/go-kratos/kratos/v2/log"
)

type SysService struct {
	pb.UnimplementedSysServer
	log                 *log.Helper
	sysAuthClient       sysV1.AuthClient
	sysAdminClient      sysV1.AdminClient
	sysRoleClient       sysV1.RoleClient
	sysPermissionClient sysV1.PermissionClient
	sysJobClient        sysV1.JobClient
	sysDeptClient       sysV1.DeptClient
	sysApiClient        sysV1.ApiClient
	sysLogClient        sysV1.LogClient
	sysDashboardClient  sysV1.DashboardClient
}

func NewSysService(
	logger log.Logger,
	sysAuthClient sysV1.AuthClient,
	sysAdminClient sysV1.AdminClient,
	sysRoleClient sysV1.RoleClient,
	sysPermissionClient sysV1.PermissionClient,
	sysJobClient sysV1.JobClient,
	sysDeptClient sysV1.DeptClient,
	sysApiClient sysV1.ApiClient,
	sysLogClient sysV1.LogClient,
	sysDashboardClient sysV1.DashboardClient,
) *SysService {
	l := log.NewHelper(log.With(logger, "module", "bff_admin/service/sys"))
	return &SysService{
		log:                 l,
		sysAuthClient:       sysAuthClient,
		sysAdminClient:      sysAdminClient,
		sysRoleClient:       sysRoleClient,
		sysPermissionClient: sysPermissionClient,
		sysJobClient:        sysJobClient,
		sysDeptClient:       sysDeptClient,
		sysApiClient:        sysApiClient,
		sysLogClient:        sysLogClient,
		sysDashboardClient:  sysDashboardClient,
	}
}

func (s *SysService) DashboardSpeech(ctx context.Context, req *pb.DashboardSpeechReq) (*pb.DashboardSpeechReply, error) {
	s.sysDashboardClient.DashboardSpeech(ctx, &sysV1.DashboardSpeechReq{})
	return &pb.DashboardSpeechReply{}, nil
}
func (s *SysService) SysAuthLoginCaptcha(ctx context.Context, req *pb.SysAuthLoginCaptchaReq) (*pb.SysAuthLoginCaptchaReply, error) {
	return &pb.SysAuthLoginCaptchaReply{}, nil
}
func (s *SysService) SysAuthLogin(ctx context.Context, req *pb.SysAuthLoginReq) (*pb.SysAuthLoginReply, error) {
	return &pb.SysAuthLoginReply{}, nil
}
func (s *SysService) SysAuthLogout(ctx context.Context, req *pb.SysAuthLogoutReq) (*pb.SysAuthLogoutReply, error) {
	return &pb.SysAuthLogoutReply{}, nil
}
func (s *SysService) SysAdminInfo(ctx context.Context, req *pb.SysAdminInfoReq) (*pb.SysAdminInfoReply, error) {
	return &pb.SysAdminInfoReply{}, nil
}
func (s *SysService) SysAdminInfoUpdate(ctx context.Context, req *pb.SysAdminInfoUpdateReq) (*pb.SysAdminInfoUpdateReply, error) {
	return &pb.SysAdminInfoUpdateReply{}, nil
}
func (s *SysService) SysAdminGenerateAvatar(ctx context.Context, req *pb.SysAdminGenerateAvatarReq) (*pb.SysAdminGenerateAvatarReply, error) {
	return &pb.SysAdminGenerateAvatarReply{}, nil
}
func (s *SysService) SysAdminPermission(ctx context.Context, req *pb.SysAdminPermissionReq) (*pb.SysAdminPermissionReply, error) {
	return &pb.SysAdminPermissionReply{}, nil
}
func (s *SysService) SysManageList(ctx context.Context, req *paginator.PaginatorReq) (*pb.SysManageListReply, error) {
	return &pb.SysManageListReply{}, nil
}
func (s *SysService) SysManageInfo(ctx context.Context, req *pb.SysManageInfoReq) (*pb.SysManageInfoReply, error) {
	return &pb.SysManageInfoReply{}, nil
}
func (s *SysService) SysManageStore(ctx context.Context, req *pb.SysManageStoreReq) (*pb.SysManageStoreReply, error) {
	return &pb.SysManageStoreReply{}, nil
}
func (s *SysService) SysManageDel(ctx context.Context, req *pb.SysManageDelReq) (*pb.SysManageDelReply, error) {
	return &pb.SysManageDelReply{}, nil
}
func (s *SysService) SysApiList(ctx context.Context, req *pb.SysApiListReq) (*pb.SysApiListReply, error) {
	return &pb.SysApiListReply{}, nil
}
func (s *SysService) SysApiStore(ctx context.Context, req *pb.SysApiStoreReq) (*pb.SysApiStoreReply, error) {
	return &pb.SysApiStoreReply{}, nil
}
func (s *SysService) SysApiDel(ctx context.Context, req *pb.SysApiDelReq) (*pb.SysApiDelReply, error) {
	return &pb.SysApiDelReply{}, nil
}
func (s *SysService) SysDeptList(ctx context.Context, req *pb.SysDeptListReq) (*pb.SysDeptListReply, error) {
	return &pb.SysDeptListReply{}, nil
}
func (s *SysService) SysDeptInfo(ctx context.Context, req *pb.SysDeptInfoReq) (*pb.SysDeptInfoReply, error) {
	return &pb.SysDeptInfoReply{}, nil
}
func (s *SysService) SysDeptStore(ctx context.Context, req *pb.SysDeptStoreReq) (*pb.SysDeptStoreReply, error) {
	return &pb.SysDeptStoreReply{}, nil
}
func (s *SysService) SysDeptDel(ctx context.Context, req *pb.SysDeptDelReq) (*pb.SysDeptDelReply, error) {
	return &pb.SysDeptDelReply{}, nil
}
func (s *SysService) SysJobList(ctx context.Context, req *paginator.PaginatorReq) (*pb.SysJobListReply, error) {
	return &pb.SysJobListReply{}, nil
}
func (s *SysService) SysJobInfo(ctx context.Context, req *pb.SysJobInfoReq) (*pb.SysJobInfoReply, error) {
	return &pb.SysJobInfoReply{}, nil
}
func (s *SysService) SysJobStore(ctx context.Context, req *pb.SysJobStoreReq) (*pb.SysJobStoreReply, error) {
	return &pb.SysJobStoreReply{}, nil
}
func (s *SysService) SysJobDel(ctx context.Context, req *pb.SysJobDelReq) (*pb.SysJobDelReply, error) {
	return &pb.SysJobDelReply{}, nil
}
func (s *SysService) SysLogList(ctx context.Context, req *paginator.PaginatorReq) (*pb.SysLogListResp, error) {
	return &pb.SysLogListResp{}, nil
}
func (s *SysService) SysLogInfo(ctx context.Context, req *pb.SysLogInfoReq) (*pb.SysLogInfoResp, error) {
	return &pb.SysLogInfoResp{}, nil
}
func (s *SysService) SysLogStore(ctx context.Context, req *pb.SysLogStoreReq) (*pb.SysLogStoreResp, error) {
	return &pb.SysLogStoreResp{}, nil
}
func (s *SysService) SysPermissionList(ctx context.Context, req *pb.SysPermissionListReq) (*pb.SysPermissionListResp, error) {
	return &pb.SysPermissionListResp{}, nil
}
func (s *SysService) SysPermissionInfo(ctx context.Context, req *pb.SysPermissionInfoReq) (*pb.SysPermissionInfoResp, error) {
	return &pb.SysPermissionInfoResp{}, nil
}
func (s *SysService) SysPermissionStore(ctx context.Context, req *pb.SysPermissionStoreReq) (*pb.SysPermissionStoreResp, error) {
	return &pb.SysPermissionStoreResp{}, nil
}
func (s *SysService) SysPermissionDel(ctx context.Context, req *pb.SysPermissionDelReq) (*pb.SysPermissionDelResp, error) {
	return &pb.SysPermissionDelResp{}, nil
}
func (s *SysService) SysPermissionStatus(ctx context.Context, req *pb.SysPermissionStatusReq) (*pb.SysPermissionStatusResp, error) {
	return &pb.SysPermissionStatusResp{}, nil
}
func (s *SysService) SysRoleList(ctx context.Context, req *pb.SysRoleListReq) (*pb.SysRoleListResp, error) {
	return &pb.SysRoleListResp{}, nil
}
func (s *SysService) SysRoleInfo(ctx context.Context, req *pb.SysRoleInfoReq) (*pb.SysRoleInfoResp, error) {
	return &pb.SysRoleInfoResp{}, nil
}
func (s *SysService) SysRoleStore(ctx context.Context, req *pb.SysRoleStoreReq) (*pb.SysRoleStoreResp, error) {
	return &pb.SysRoleStoreResp{}, nil
}
func (s *SysService) SysRoleDel(ctx context.Context, req *pb.SysRoleDelReq) (*pb.SysRoleDelResp, error) {
	return &pb.SysRoleDelResp{}, nil
}
