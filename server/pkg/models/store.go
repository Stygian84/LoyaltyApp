package models

import(
  "context"
  "database/sql"
  "fmt"
  "time"
)
type TransferParams struct{
  UserId int32 `json:"userId" binding:"required"` 
  CreditToTransfer float64 `json:"creditToTransfer" binding:"required"`
  MembershipId string `json:"membershipId" binding:"required"`
  ProgramId int32 `json:"programId" binding:"required"`
  RewardShouldReceive float64 `json:"rewardShouldReceive"`
}
type Store struct{
  *Queries
  db *sql.DB
}

func NewStore(db *sql.DB) *Store{
  return &Store{
    db:db,
    Queries:New(db),
  }
}

func (store *Store) execTx(ctx context.Context,fn func(*Queries) error) error{
  tx,err := store.db.BeginTx(ctx,nil)
  if err!=nil{
    return err
  }
  q := New(tx)
  err = fn(q)
  if err!=nil{
    if rbErr:=tx.Rollback(); rbErr!=nil{
      return fmt.Errorf("tx err: %v, rb err: %v",err,rbErr)
    }
    return err
  }
  return tx.Commit()
}

func (store *Store) CreditTransferOut(ctx context.Context,arg TransferParams,promo sql.NullInt32) (CreditRequest,error){
  var result CreditRequest
  err:= store.execTx(ctx, func (q *Queries)error{
    var err error
    request:=CreateCreditRequestParams{
      UserID:arg.UserId,
      Program:arg.ProgramId,
      MemberID:arg.MembershipId,
      CreditUsed:arg.CreditToTransfer,
      RewardShouldReceive:arg.RewardShouldReceive,
      PromoUsed:promo,
      TransactionTime:sql.NullTime{Time:time.Now(),Valid:true},
      TransactionStatus:TransactionStatusEnum("created"),
    }
    result,err =q.CreateCreditRequest(
     ctx,request, 
    )
    if err!=nil{
      return err
    }
    balanceParam := DecreBalanceParams {
      CreditBalance:arg.CreditToTransfer,
      ID: int64(arg.UserId),
    }

    //TODO prevent deadlock
    err = q.DecreBalance(ctx,balanceParam)
    if err!=nil{
      return err
    }
    return nil
      
  })
  return result,err 

}
