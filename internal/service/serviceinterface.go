package service

import (
	pb "github.com/ratheeshkumar25/ultimate-rent-car/ultimate-rent-user-service/internal/pb/ulimateuser"
)

// ...existing code...
type UserServiceInter interface {
	//User Authentication Services
	SignupService(p *pb.SignupRequest) (*pb.SignupResponse, error)
	VerifyService(p *pb.OtpRequest) (*pb.Response, error)
	LoginService(p *pb.LoginRequest) (*pb.LoginResponse, error)
	ChangePassword(p *pb.ChangePasswordRequest) (*pb.Response, error)

	//Document Management
	UploadDocumentService(p *pb.UploadDocumentRequest) (*pb.Response, error)
	GetDocumentsService(p *pb.ID) (*pb.Response, error)
	DeleteDocumentService(p *pb.ID) (*pb.Response, error)

	//Profile Managemwnr
	GetProfileService(p *pb.ID) (*pb.ProfileResponse, error)
	UpdateProfileService(p *pb.UpdateProfileRequest) (*pb.Response, error)

	//Address Management
	AddAddressService(p *pb.AddAddressRequest) (*pb.Response, error)
	UpdateAddressService(p *pb.UpdateAddressRequest) (*pb.Response, error)
	DeleteAddressService(p *pb.ID) (*pb.Response, error)
	ListAddressesService(p *pb.ID) (*pb.Response, error)

	//Wallet Management
	GetWalletBalanceService(p *pb.ID) (*pb.WalletResponse, error)
	UpdateWallentService(p *pb.UpdateWalletRequest) (*pb.Response, error)

	//Admin Operations Services
	BlockUserService(p *pb.ID) (*pb.Response, error)
	UnblockUserService(p *pb.ID) (*pb.Response, error)
	GetAllUsersService(p *pb.NoParam) (*pb.Response, error)
}
