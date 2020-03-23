package zoracle

import (
	"github.com/cosmos/gaia/x/zoracle/internal/keeper"
	"github.com/cosmos/gaia/x/zoracle/internal/types"
)

const (
	ModuleName        = types.ModuleName
	DefaultParamspace = types.DefaultParamspace
	RouterKey         = types.RouterKey
	StoreKey          = types.StoreKey

	EventTypeCreateDataSource   = types.EventTypeCreateDataSource
	EventTypeEditDataSource     = types.EventTypeEditDataSource
	EventTypeCreateOracleScript = types.EventTypeCreateOracleScript
	EventTypeEditOracleScript   = types.EventTypeEditOracleScript
	EventTypeRequest            = types.EventTypeRequest
	EventTypeReport             = types.EventTypeReport

	AttributeKeyID        = types.AttributeKeyID
	AttributeKeyRequestID = types.AttributeKeyRequestID
	AttributeKeyValidator = types.AttributeKeyValidator
)

var (
	NewKeeper                  = keeper.NewKeeper
	NewQuerier                 = keeper.NewQuerier
	ModuleCdc                  = types.ModuleCdc
	RegisterCodec              = types.RegisterCodec
	NewMsgRequestData          = types.NewMsgRequestData
	NewMsgReportData           = types.NewMsgReportData
	NewMsgCreateOracleScript   = types.NewMsgCreateOracleScript
	NewMsgEditOracleScript     = types.NewMsgEditOracleScript
	NewMsgCreateDataSource     = types.NewMsgCreateDataSource
	NewMsgEditDataSource       = types.NewMsgEditDataSource
	NewMsgAddOracleAddress     = types.NewMsgAddOracleAddress
	NewMsgRemoveOracleAdderess = types.NewMsgRemoveOracleAdderess
	NewOraclePacketData        = types.NewOraclePacketData

	RequestStoreKey      = types.RequestStoreKey
	ResultStoreKey       = types.ResultStoreKey
	DataSourceStoreKey   = types.DataSourceStoreKey
	OracleScriptStoreKey = types.OracleScriptStoreKey

	NewParams              = types.NewParams
	NewDataSource          = types.NewDataSource
	NewOracleScript        = types.NewOracleScript
	DefaultParams          = types.DefaultParams
	NewRawDataReport       = types.NewRawDataReport
	NewRawDataReportWithID = types.NewRawDataReportWithID

	KeyMaxDataSourceExecutableSize  = types.KeyMaxDataSourceExecutableSize
	KeyMaxOracleScriptCodeSize      = types.KeyMaxOracleScriptCodeSize
	KeyMaxCalldataSize              = types.KeyMaxCalldataSize
	KeyMaxDataSourceCountPerRequest = types.KeyMaxDataSourceCountPerRequest
	KeyMaxRawDataReportSize         = types.KeyMaxRawDataReportSize
	KeyMaxResultSize                = types.KeyMaxResultSize

	QueryRequestByID    = types.QueryRequestByID
	QueryRequests       = types.QueryRequests
	QueryPending        = types.QueryPending
	QueryRequestNumber  = types.QueryRequestNumber
	QueryDataSourceByID = types.QueryDataSourceByID
	QueryDataSources    = types.QueryDataSources
	QueryOracleScripts  = types.QueryOracleScripts

	ParamKeyTable = keeper.ParamKeyTable
)

type (
	Keeper                  = keeper.Keeper
	MsgRequestData          = types.MsgRequestData
	MsgReportData           = types.MsgReportData
	MsgCreateDataSource     = types.MsgCreateDataSource
	MsgEditDataSource       = types.MsgEditDataSource
	MsgCreateOracleScript   = types.MsgCreateOracleScript
	MsgEditOracleScript     = types.MsgEditOracleScript
	MsgAddOracleAddress     = types.MsgAddOracleAddress
	MsgRemoveOracleAdderess = types.MsgRemoveOracleAdderess
	OraclePacketData        = types.OraclePacketData

	RawDataReport         = types.RawDataReport
	RawDataReportWithID   = types.RawDataReportWithID
	RequestQuerierInfo    = types.RequestQuerierInfo
	DataSourceQuerierInfo = types.DataSourceQuerierInfo

	RequestID      = types.RequestID
	OracleScriptID = types.OracleScriptID
	ExternalID     = types.ExternalID
	DataSourceID   = types.DataSourceID

	DataSource   = types.DataSource
	OracleScript = types.OracleScript
)
