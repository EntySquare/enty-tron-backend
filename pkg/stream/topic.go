package stream

type Topic string

// scenes: DEPOSIT / TRANSFER / GEN_PAY_ORDER / PAY / WITHDRAW

// Topic case DEPOSIT:
// -
// General
const ORDER_RESP Topic = "ORDER_RESP"

// -
// <GateHandler>: x
// <OMS>/<ChainHandler>: DEPOSIT_ORDER
const DEPOSIT_ORDER Topic = "DEPOSIT_ORDER"

// Topic case TRANSFER:
// -
// <GateHandler>/<OMS>: TRANSFER_ORDER ORDER_RESP
const TRANSFER_ORDER Topic = "TRANSFER_ORDER"

// Topic case GEN_PAY_ORDER:
// -
// <GateHandler>/<OMS>: MAKE_ORDER ORDER_RESP
const MAKE_ORDER Topic = "MAKE_ORDER"

// Topic case GEN_PAY_ORDER:
// -
// <GateHandler>/<OMS>: PAY_ORDER ORDER_RESP
const PAY_ORDER Topic = "PAY_ORDER"

// Topic case WITHDRAW:
// -
// <GateHandler>/<OMS>: WITHDRAW_ORDER ORDER_RESP
const WITHDRAW_ORDER Topic = "WITHDRAW_ORDER"

// <OMS>/<ChainHandler>: TRANSACTION TRANSACTION_RESP TRANSACTION_UPDATE
const TRANSACTION Topic = "TRANSACTION"
const TRANSACTION_RESP Topic = "TRANSACTION_RESP"
const TRANSACTION_UPDATE Topic = "TRANSACTION_UPDATE"
const WITHDRAW_COLLECT Topic = "WITHDRAW_COLLECT"
const RISK_REQUEST Topic = "RISK_REQUEST"
const RISK_RESPONSE Topic = "RISK_RESPONSE"
const COLLECT_CHECK Topic = "COLLECT_CHECK"
const CHAIN_PANIC Topic = "CHAIN_PANIC"
