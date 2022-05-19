package tron

//web status
const UNCOMMITTED = "0"
const WAITING_FOR_PAYMENT = "1"
const WAITING_FOR_DELIVERY = "2"
const DELIVERED = "3"

//ret code
const SUCCESS = "0"
const SOLD = "1"
const OVERLIMIT = "1"
const IDOOVERLIMIT = "3"
const LOCKED = "2"

//limit const
const TBLIMIT = 2
const NFTLIMIT = 2

//txs status
const TXS_WAITING_FOR_CHAIN_CONFIRM = "0"
const TXS_CHAIN_CONFIRM = "1"
const TXS_CHAIN_FAILED = "-1"
const TXS_ADMIN_CONFIRM = "2"

//Txs type
const TYPE_TB = "1"
const TYPE_NFT = "2"
