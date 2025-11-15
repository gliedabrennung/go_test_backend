package repo

import (
	"time"

	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type Config struct {
	SkipDefaultTransaction    bool
	DefaultTransactionTimeout time.Duration
	DefaultContextTimeout     time.Duration

	NamingStrategy       schema.Namer
	FullSaveAssociations bool
	Logger               logger.Interface
	NowFunc              func() time.Time
	DryRun               bool
	PrepareStmt          bool
	PrepareStmtMaxSize   int
	PrepareStmtTTL       time.Duration

	DisableAutomaticPing                     bool
	DisableForeignKeyConstraintWhenMigrating bool
	IgnoreRelationshipsWhenMigrating         bool
	DisableNestedTransaction                 bool
	AllowGlobalUpdate                        bool
	QueryFields                              bool
	CreateBatchSize                          int
	TranslateError                           bool
	PropagateUnscoped                        bool
}
