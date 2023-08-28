package service

import (
	"context"
	pb "fkratos/api/bff_admin/v1"
	"fkratos/api/paginator"
	sysV1 "fkratos/api/rpc_sys/v1"
	"fkratos/internal/errorx"
	"fkratos/internal/meta"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
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
	sysAPIClient        sysV1.APIClient
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
	sysAPIClient sysV1.APIClient,
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
		sysAPIClient:        sysAPIClient,
		sysLogClient:        sysLogClient,
		sysDashboardClient:  sysDashboardClient,
	}
}

func (s *SysService) DashboardSpeech(ctx context.Context, _ *pb.DashboardSpeechReq) (*pb.DashboardSpeechReply, error) {
	resp := new(pb.DashboardSpeechReply)
	result, err := s.sysDashboardClient.DashboardSpeech(ctx, &sysV1.DashboardSpeechReq{})
	if err != nil {
		return nil, err
	}
	err = copier.Copy(resp, result)
	if err != nil {
		return nil, errorx.DataFormattingError.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	return resp, nil
}
func (s *SysService) SysAuthLoginCaptcha(ctx context.Context, _ *pb.SysAuthLoginCaptchaReq) (*pb.SysAuthLoginCaptchaReply, error) {
	resp := new(pb.SysAuthLoginCaptchaReply)
	result, err := s.sysAuthClient.SysAuthLoginCaptcha(ctx, &sysV1.SysAuthLoginCaptchaReq{})
	if err != nil {
		return nil, err
	}
	err = copier.Copy(resp, result)
	if err != nil {
		return nil, errorx.DataFormattingError.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	return resp, nil
}
func (s *SysService) SysAuthLogin(ctx context.Context, req *pb.SysAuthLoginReq) (*pb.SysAuthLoginReply, error) {
	resp := new(pb.SysAuthLoginReply)
	result, err := s.sysAuthClient.SysAuthLogin(ctx, &sysV1.SysAuthLoginReq{
		CaptchaId:  req.GetCaptchaId(),
		VerifyCode: req.GetVerifyCode(),
		Username:   req.GetUsername(),
		Password:   req.GetPassword(),
	})
	if err != nil {
		return nil, err
	}
	err = copier.Copy(resp, result)
	if err != nil {
		return nil, errorx.DataFormattingError.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	return resp, nil
}
func (s *SysService) SysAuthLogout(ctx context.Context, _ *pb.SysAuthLogoutReq) (*pb.SysAuthLogoutReply, error) {
	resp := new(pb.SysAuthLogoutReply)
	result, err := s.sysAuthClient.SysAuthLogout(ctx, &sysV1.SysAuthLogoutReq{
		AdminId: meta.GetAdminIDFromClient(ctx),
	})
	if err != nil {
		return nil, err
	}
	err = copier.Copy(resp, result)
	if err != nil {
		return nil, errorx.DataFormattingError.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	return resp, nil
}
func (s *SysService) SysAdminInfo(ctx context.Context, _ *pb.SysAdminInfoReq) (*pb.SysAdminInfoReply, error) {
	resp := new(pb.SysAdminInfoReply)
	result, err := s.sysAdminClient.SysAdminInfo(ctx, &sysV1.SysAdminInfoReq{
		AdminId: meta.GetAdminIDFromClient(ctx),
	})
	if err != nil {
		return nil, err
	}
	err = copier.Copy(resp, result)
	if err != nil {
		return nil, errorx.DataFormattingError.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	return resp, nil
}
func (s *SysService) SysAdminInfoUpdate(ctx context.Context, req *pb.SysAdminInfoUpdateReq) (*pb.SysAdminInfoUpdateReply, error) {
	resp := new(pb.SysAdminInfoUpdateReply)
	result, err := s.sysAdminClient.SysAdminInfoUpdate(ctx, &sysV1.SysAdminInfoUpdateReq{
		AdminId:  meta.GetAdminIDFromClient(ctx),
		Nickname: req.GetNickname(),
		Email:    req.GetEmail(),
		Mobile:   req.GetMobile(),
		Motto:    req.GetMotto(),
		Password: req.GetPassword(),
	})
	if err != nil {
		return nil, err
	}
	err = copier.Copy(resp, result)
	if err != nil {
		return nil, errorx.DataFormattingError.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	return resp, nil
}
func (s *SysService) SysAdminGenerateAvatar(ctx context.Context, _ *pb.SysAdminGenerateAvatarReq) (*pb.SysAdminGenerateAvatarReply, error) {
	resp := new(pb.SysAdminGenerateAvatarReply)
	result, err := s.sysAdminClient.SysAdminGenerateAvatar(ctx, &sysV1.SysAdminGenerateAvatarReq{
		AdminId: meta.GetAdminIDFromClient(ctx),
	})
	if err != nil {
		return nil, err
	}
	err = copier.Copy(resp, result)
	if err != nil {
		return nil, errorx.DataFormattingError.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	return resp, nil
}
func (s *SysService) SysAdminPermission(ctx context.Context, _ *pb.SysAdminPermissionReq) (*pb.SysAdminPermissionReply, error) {
	resp := new(pb.SysAdminPermissionReply)
	result, err := s.sysAdminClient.SysAdminPermission(ctx, &sysV1.SysAdminPermissionReq{
		AdminId: meta.GetAdminIDFromClient(ctx),
	})
	if err != nil {
		return nil, err
	}
	err = copier.Copy(resp, result)
	if err != nil {
		return nil, errorx.DataFormattingError.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	return resp, nil
}
func (s *SysService) SysManageList(ctx context.Context, req *paginator.PaginatorReq) (*pb.SysManageListReply, error) {
	resp := new(pb.SysManageListReply)
	result, err := s.sysAdminClient.SysManageList(ctx, &paginator.PaginatorReq{
		Page:        req.GetPage(),
		PageSize:    req.GetPageSize(),
		QuickSearch: req.GetQuickSearch(),
		Search:      req.GetSearch(),
		Order:       req.GetOrder(),
	})
	if err != nil {
		return nil, err
	}
	err = copier.Copy(resp, result)
	if err != nil {
		return nil, errorx.DataFormattingError.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	return resp, nil
}
func (s *SysService) SysManageInfo(ctx context.Context, req *pb.SysManageInfoReq) (*pb.SysManageInfoReply, error) {
	resp := new(pb.SysManageInfoReply)
	result, err := s.sysAdminClient.SysManageInfo(ctx, &sysV1.SysManageInfoReq{
		AdminId: req.GetAdminId(),
	})
	if err != nil {
		return nil, err
	}
	err = copier.Copy(resp, result)
	if err != nil {
		return nil, errorx.DataFormattingError.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	return resp, nil
}
func (s *SysService) SysManageStore(ctx context.Context, req *pb.SysManageStoreReq) (*pb.SysManageStoreReply, error) {
	resp := new(pb.SysManageStoreReply)
	result, err := s.sysAdminClient.SysManageStore(ctx, &sysV1.SysManageStoreReq{
		Id:       req.GetId(),
		Username: req.GetUsername(),
		Nickname: req.GetNickname(),
		Password: req.GetPassword(),
		Avatar:   req.GetAvatar(),
		Gender:   req.GetGender(),
		Email:    req.GetEmail(),
		Mobile:   req.GetMobile(),
		JobID:    req.GetJobID(),
		DeptID:   req.GetDeptID(),
		RoleIds:  req.GetRoleIds(),
		Motto:    req.GetMotto(),
		Status:   req.GetStatus(),
	})
	if err != nil {
		return nil, err
	}
	err = copier.Copy(resp, result)
	if err != nil {
		return nil, errorx.DataFormattingError.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	return resp, nil
}
func (s *SysService) SysManageDel(ctx context.Context, req *pb.SysManageDelReq) (*pb.SysManageDelReply, error) {
	resp := new(pb.SysManageDelReply)
	result, err := s.sysAdminClient.SysManageDel(ctx, &sysV1.SysManageDelReq{
		Ids: req.GetIds(),
	})
	if err != nil {
		return nil, err
	}
	err = copier.Copy(resp, result)
	if err != nil {
		return nil, errorx.DataFormattingError.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	return resp, nil
}
func (s *SysService) SysAPIList(ctx context.Context, req *pb.SysAPIListReq) (*pb.SysAPIListReply, error) {
	resp := new(pb.SysAPIListReply)
	result, err := s.sysAPIClient.SysAPIList(ctx, &sysV1.SysAPIListReq{
		PermissionId: req.GetPermissionId(),
	})
	if err != nil {
		return nil, err
	}
	err = copier.Copy(resp, result)
	if err != nil {
		return nil, errorx.DataFormattingError.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	return resp, nil
}
func (s *SysService) SysAPIStore(ctx context.Context, req *pb.SysAPIStoreReq) (*pb.SysAPIStoreReply, error) {
	resp := new(pb.SysAPIStoreReply)
	result, err := s.sysAPIClient.SysAPIStore(ctx, &sysV1.SysAPIStoreReq{
		Id:           req.GetId(),
		PermissionID: req.GetPermissionID(),
		Method:       req.GetMethod(),
		Path:         req.GetPath(),
		Desc:         req.GetDesc(),
	})
	if err != nil {
		return nil, err
	}
	err = copier.Copy(resp, result)
	if err != nil {
		return nil, errorx.DataFormattingError.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	return resp, nil
}
func (s *SysService) SysAPIDel(ctx context.Context, req *pb.SysAPIDelReq) (*pb.SysAPIDelReply, error) {
	resp := new(pb.SysAPIDelReply)
	result, err := s.sysAPIClient.SysAPIDel(ctx, &sysV1.SysAPIDelReq{
		Ids: req.GetIds(),
	})
	if err != nil {
		return nil, err
	}
	err = copier.Copy(resp, result)
	if err != nil {
		return nil, errorx.DataFormattingError.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	return resp, nil
}
func (s *SysService) SysDeptList(ctx context.Context, _ *pb.SysDeptListReq) (*pb.SysDeptListReply, error) {
	resp := new(pb.SysDeptListReply)
	result, err := s.sysDeptClient.SysDeptList(ctx, &sysV1.SysDeptListReq{})
	if err != nil {
		return nil, err
	}
	err = copier.Copy(resp, result)
	if err != nil {
		return nil, errorx.DataFormattingError.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	return resp, nil
}
func (s *SysService) SysDeptInfo(ctx context.Context, req *pb.SysDeptInfoReq) (*pb.SysDeptInfoReply, error) {
	resp := new(pb.SysDeptInfoReply)
	result, err := s.sysDeptClient.SysDeptInfo(ctx, &sysV1.SysDeptInfoReq{
		Id: req.GetId(),
	})
	if err != nil {
		return nil, err
	}
	err = copier.Copy(resp, result)
	if err != nil {
		return nil, errorx.DataFormattingError.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	return resp, nil
}
func (s *SysService) SysDeptStore(ctx context.Context, req *pb.SysDeptStoreReq) (*pb.SysDeptStoreReply, error) {
	resp := new(pb.SysDeptStoreReply)
	result, err := s.sysDeptClient.SysDeptStore(ctx, &sysV1.SysDeptStoreReq{
		Id:          req.GetId(),
		Pid:         req.GetPid(),
		Name:        req.GetName(),
		FullName:    req.GetFullName(),
		Responsible: req.GetResponsible(),
		Phone:       req.GetPhone(),
		Email:       req.GetEmail(),
		Type:        req.GetType(),
		Status:      req.GetStatus(),
		Sort:        req.GetSort(),
	})
	if err != nil {
		return nil, err
	}
	err = copier.Copy(resp, result)
	if err != nil {
		return nil, errorx.DataFormattingError.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	return resp, nil
}
func (s *SysService) SysDeptDel(ctx context.Context, req *pb.SysDeptDelReq) (*pb.SysDeptDelReply, error) {
	resp := new(pb.SysDeptDelReply)
	result, err := s.sysDeptClient.SysDeptDel(ctx, &sysV1.SysDeptDelReq{
		Ids: req.GetIds(),
	})
	if err != nil {
		return nil, err
	}
	err = copier.Copy(resp, result)
	if err != nil {
		return nil, errorx.DataFormattingError.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	return resp, nil
}
func (s *SysService) SysJobList(ctx context.Context, req *paginator.PaginatorReq) (*pb.SysJobListReply, error) {
	resp := new(pb.SysJobListReply)
	result, err := s.sysJobClient.SysJobList(ctx, &paginator.PaginatorReq{
		Page:        req.GetPage(),
		PageSize:    req.GetPageSize(),
		QuickSearch: req.GetQuickSearch(),
		Search:      req.GetSearch(),
		Order:       req.GetOrder(),
	})
	if err != nil {
		return nil, err
	}
	err = copier.Copy(resp, result)
	if err != nil {
		return nil, errorx.DataFormattingError.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	return resp, nil
}
func (s *SysService) SysJobInfo(ctx context.Context, req *pb.SysJobInfoReq) (*pb.SysJobInfoReply, error) {
	resp := new(pb.SysJobInfoReply)
	result, err := s.sysJobClient.SysJobInfo(ctx, &sysV1.SysJobInfoReq{
		Id: req.GetId(),
	})
	if err != nil {
		return nil, err
	}
	err = copier.Copy(resp, result)
	if err != nil {
		return nil, errorx.DataFormattingError.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	return resp, nil
}
func (s *SysService) SysJobStore(ctx context.Context, req *pb.SysJobStoreReq) (*pb.SysJobStoreReply, error) {
	resp := new(pb.SysJobStoreReply)
	result, err := s.sysJobClient.SysJobStore(ctx, &sysV1.SysJobStoreReq{
		Id:     req.GetId(),
		Name:   req.GetName(),
		Code:   req.GetCode(),
		Remark: req.GetRemark(),
		Status: req.GetStatus(),
		Sort:   req.GetSort(),
	})
	if err != nil {
		return nil, err
	}
	err = copier.Copy(resp, result)
	if err != nil {
		return nil, errorx.DataFormattingError.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	return resp, nil
}
func (s *SysService) SysJobDel(ctx context.Context, req *pb.SysJobDelReq) (*pb.SysJobDelReply, error) {
	resp := new(pb.SysJobDelReply)
	result, err := s.sysJobClient.SysJobDel(ctx, &sysV1.SysJobDelReq{
		Ids: req.GetIds(),
	})
	if err != nil {
		return nil, err
	}
	err = copier.Copy(resp, result)
	if err != nil {
		return nil, errorx.DataFormattingError.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	return resp, nil
}
func (s *SysService) SysLogList(ctx context.Context, req *paginator.PaginatorReq) (*pb.SysLogListResp, error) {
	resp := new(pb.SysLogListResp)
	result, err := s.sysLogClient.SysLogList(ctx, &paginator.PaginatorReq{
		Page:        req.GetPage(),
		PageSize:    req.GetPageSize(),
		QuickSearch: req.GetQuickSearch(),
		Search:      req.GetSearch(),
		Order:       req.GetOrder(),
	})
	if err != nil {
		return nil, err
	}
	err = copier.Copy(resp, result)
	if err != nil {
		return nil, errorx.DataFormattingError.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	return resp, nil
}
func (s *SysService) SysLogInfo(ctx context.Context, req *pb.SysLogInfoReq) (*pb.SysLogInfoResp, error) {
	resp := new(pb.SysLogInfoResp)
	result, err := s.sysLogClient.SysLogInfo(ctx, &sysV1.SysLogInfoReq{
		Id: req.GetId(),
	})
	if err != nil {
		return nil, err
	}
	err = copier.Copy(resp, result)
	if err != nil {
		return nil, errorx.DataFormattingError.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	return resp, nil
}
func (s *SysService) SysLogStore(ctx context.Context, req *pb.SysLogStoreReq) (*pb.SysLogStoreResp, error) {
	resp := new(pb.SysLogStoreResp)
	result, err := s.sysLogClient.SysLogStore(ctx, &sysV1.SysLogStoreReq{
		AdminID:   req.GetAdminID(),
		Username:  req.GetUsername(),
		Ip:        req.GetIp(),
		Uri:       req.GetUri(),
		UriDesc:   req.GetUriDesc(),
		Useragent: req.GetUseragent(),
		Req:       req.GetReq(),
		Resp:      req.GetResp(),
	})
	if err != nil {
		return nil, err
	}
	err = copier.Copy(resp, result)
	if err != nil {
		return nil, errorx.DataFormattingError.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	return resp, nil
}
func (s *SysService) SysPermissionList(ctx context.Context, _ *pb.SysPermissionListReq) (*pb.SysPermissionListResp, error) {
	resp := new(pb.SysPermissionListResp)
	result, err := s.sysPermissionClient.SysPermissionList(ctx, &sysV1.SysPermissionListReq{})
	if err != nil {
		return nil, err
	}
	err = copier.Copy(resp, result)
	if err != nil {
		return nil, errorx.DataFormattingError.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	return resp, nil
}
func (s *SysService) SysPermissionInfo(ctx context.Context, req *pb.SysPermissionInfoReq) (*pb.SysPermissionInfoResp, error) {
	resp := new(pb.SysPermissionInfoResp)
	result, err := s.sysPermissionClient.SysPermissionInfo(ctx, &sysV1.SysPermissionInfoReq{
		Id: req.GetId(),
	})
	if err != nil {
		return nil, err
	}
	err = copier.Copy(resp, result)
	if err != nil {
		return nil, errorx.DataFormattingError.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	return resp, nil
}
func (s *SysService) SysPermissionStore(ctx context.Context, req *pb.SysPermissionStoreReq) (*pb.SysPermissionStoreResp, error) {
	resp := new(pb.SysPermissionStoreResp)
	result, err := s.sysPermissionClient.SysPermissionStore(ctx, &sysV1.SysPermissionStoreReq{
		Id:        req.GetId(),
		Pid:       req.GetPid(),
		Type:      req.GetType(),
		Title:     req.GetTitle(),
		Name:      req.GetName(),
		Path:      req.GetPath(),
		Icon:      req.GetIcon(),
		MenuType:  req.GetMenuType(),
		Url:       req.GetUrl(),
		Component: req.GetComponent(),
		Keepalive: req.GetKeepalive(),
		Extend:    req.GetExtend(),
		Remark:    req.GetRemark(),
		Sort:      req.GetSort(),
		Status:    req.GetStatus(),
	})
	if err != nil {
		return nil, err
	}
	err = copier.Copy(resp, result)
	if err != nil {
		return nil, errorx.DataFormattingError.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	return resp, nil
}
func (s *SysService) SysPermissionDel(ctx context.Context, req *pb.SysPermissionDelReq) (*pb.SysPermissionDelResp, error) {
	resp := new(pb.SysPermissionDelResp)
	result, err := s.sysPermissionClient.SysPermissionDel(ctx, &sysV1.SysPermissionDelReq{
		Ids: req.GetIds(),
	})
	if err != nil {
		return nil, err
	}
	err = copier.Copy(resp, result)
	if err != nil {
		return nil, errorx.DataFormattingError.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	return resp, nil
}
func (s *SysService) SysPermissionStatus(ctx context.Context, req *pb.SysPermissionStatusReq) (*pb.SysPermissionStatusResp, error) {
	resp := new(pb.SysPermissionStatusResp)
	result, err := s.sysPermissionClient.SysPermissionStatus(ctx, &sysV1.SysPermissionStatusReq{
		Id:     req.GetId(),
		Status: req.GetStatus(),
	})
	if err != nil {
		return nil, err
	}
	err = copier.Copy(resp, result)
	if err != nil {
		return nil, errorx.DataFormattingError.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	return resp, nil
}
func (s *SysService) SysRoleList(ctx context.Context, _ *pb.SysRoleListReq) (*pb.SysRoleListResp, error) {
	resp := new(pb.SysRoleListResp)
	result, err := s.sysRoleClient.SysRoleList(ctx, &sysV1.SysRoleListReq{})
	if err != nil {
		return nil, err
	}
	err = copier.Copy(resp, result)
	if err != nil {
		return nil, errorx.DataFormattingError.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	return resp, nil
}
func (s *SysService) SysRoleInfo(ctx context.Context, req *pb.SysRoleInfoReq) (*pb.SysRoleInfoResp, error) {
	resp := new(pb.SysRoleInfoResp)
	result, err := s.sysRoleClient.SysRoleInfo(ctx, &sysV1.SysRoleInfoReq{
		Id: req.GetId(),
	})
	if err != nil {
		return nil, err
	}
	err = copier.Copy(resp, result)
	if err != nil {
		return nil, errorx.DataFormattingError.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	return resp, nil
}
func (s *SysService) SysRoleStore(ctx context.Context, req *pb.SysRoleStoreReq) (*pb.SysRoleStoreResp, error) {
	resp := new(pb.SysRoleStoreResp)
	result, err := s.sysRoleClient.SysRoleStore(ctx, &sysV1.SysRoleStoreReq{
		Id:            req.GetId(),
		Pid:           req.GetPid(),
		Name:          req.GetName(),
		Remark:        req.GetRemark(),
		Status:        req.GetStatus(),
		PermissionIds: req.GetPermissionIds(),
	})
	if err != nil {
		return nil, err
	}
	err = copier.Copy(resp, result)
	if err != nil {
		return nil, errorx.DataFormattingError.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	return resp, nil
}
func (s *SysService) SysRoleDel(ctx context.Context, req *pb.SysRoleDelReq) (*pb.SysRoleDelResp, error) {
	resp := new(pb.SysRoleDelResp)
	result, err := s.sysRoleClient.SysRoleDel(ctx, &sysV1.SysRoleDelReq{
		Ids: req.GetIds(),
	})
	if err != nil {
		return nil, err
	}
	err = copier.Copy(resp, result)
	if err != nil {
		return nil, errorx.DataFormattingError.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	return resp, nil
}
