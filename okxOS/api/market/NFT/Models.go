package nft

type NFTModel struct {
	Name              string             `json:"name"`
	TokenId           string             `json:"tokenId"`
	TokenUrl          string             `json:"tokenUrl"`
	Image             string             `json:"image"`
	ImagePreviewUrl   string             `json:"imagePreviewUrl"`
	ImageThumbnailUrl string             `json:"imageThumbnailUrl"`
	AnimationUrl      string             `json:"animationUrl"` // OKX的响应是string
	Attributes        string             `json:"attributes"`
	AssetContract     AssetContractModel `json:"assetContract"`
	Collection        CollectionModel    `json:"collection"`
}

type AttributesModel struct {
	TraitType   string `json:"trait_type"`
	Value       string `json:"value"`
	DisplayType string `json:"display_type"`
} // NFT属性模型

type NFTAssetsModel struct {
	Name              string             `json:"name"`
	TokenId           string             `json:"tokenId"`
	Amount            string             `json:"amount"`
	TokenUrl          string             `json:"tokenUrl"`
	Image             string             `json:"image"`
	ImagePreviewUrl   string             `json:"imagePreviewUrl"`
	ImageThumbnailUrl string             `json:"imageThumbnailUrl"`
	AnimationUrl      string             `json:"animationUrl"`
	Attributes        string             `json:"attributes"` // OKX的响应是string
	AssetContract     AssetContractModel `json:"assetContract"`
	Collection        CollectionModel    `json:"collection"`
	OwnerAddress      string             `json:"ownerAddress"`
	IsLazyMintType    bool               `json:"isLazyMintType"`
} // NFT资产模型

type CollectionModel struct {
	Name            string               `json:"name"`            // 合集名称
	Des             string               `json:"des"`             // 合集简介
	Image           string               `json:"image"`           // 合集头像图片地址
	BackgroundImage string               `json:"backgroundImage"` // 合集背景图片地址
	Slug            string               `json:"slug"`            // 合集 slug 名称，即合集的唯一标识
	CertificateFlag bool                 `json:"certificateFlag"` // 合集认证标记
	OfficialWebsite string               `json:"officialWebsite"` // 合集官网地址
	InstagramUrl    string               `json:"instagramUrl"`    // 合集 Instagram 地址
	DiscordUrl      string               `json:"discordUrl"`      // 合集 Discord 地址
	MediumUrl       string               `json:"mediumUrl"`       // 合集 Medium 地址
	TwitterUrl      string               `json:"twitterUrl"`      // 合集推特地址
	CategoryList    []string             `json:"categoryList"`    // 合集分类标签列表
	AssetContracts  []AssetContractModel `json:"assetContracts"`  // 与该合集相关的合约信息列表，包括合约地址，协议标准，拥有者地址，是否支持 erc2981 协议等，一个合约信息模型数组
	Stats           CollectionStatsModel `json:"stats"`           // 该合集对应统计数据，包括交易额、地板价等，一个合集统计信息模型对象
	IsHidden        bool                 `json:"isHidden"`        // NFTDetails不返回，NFTList返回
} // 合集模型

type AssetContractModel struct {
	Chain           string `json:"chain"`           // 与该合集相关的合约所属链
	ContractAddress string `json:"contractAddress"` // 与该合集相关的合约地址
	TokenStandard   string `json:"tokenStandard"`   // 该合约对应的协议类型，例如 ERC721、ERC1155
	OwnerAddress    string `json:"ownerAddress"`    // 该合约对应的拥有者地址
	Erc2981         bool   `json:"erc2981"`         // 该合约是否支持 2981 协议
} // 合约信息模型

type CollectionStatsModel struct {
	LatestPrice string `json:"latestPrice"` // 该合集的最新成交价格
	TotalVolume string `json:"totalVolume"` // 该合集的总交易额
	TotalCount  string `json:"totalCount"`  // 该合集的NFT数量
	OwnerCount  string `json:"ownerCount"`  // 该合集的拥有者数量
	FloorPrice  string `json:"floorPrice"`  // 该合集的地板价
} // 合约统计信息模型

// 订单状态枚举
// active	表示该订单当前有效，可以被成交
// inactive	表示该订单当前暂时失效，在后续可能重新变为 active 状态
// cancelled	表示该订单已经被取消，无法被交易
// sold	表示该订单已经被成交，无法再次被交易

type OKXOrdersModel struct {
	OrderId           string            `json:"orderId"`
	CreateTime        int64             `json:"createTime"`
	UpdateTime        int64             `json:"updateTime"`
	ListingTime       int64             `json:"listingTime"`
	ExpirationTime    int64             `json:"expirationTime"`
	Status            string            `json:"status"`
	OrderHash         string            `json:"orderHash"`
	ProtocolData      ProtocolDataModel `json:"protocolData"` // 订单参数
	ProtocolAddress   string            `json:"protocolAddress"`
	Chain             string            `json:"chain"`
	Maker             string            `json:"maker"`
	OrderType         string            `json:"orderType"`
	Price             string            `json:"price"`
	CurrencyAddress   string            `json:"currencyAddress"`
	CollectionAddress string            `json:"collectionAddress"`
	TokenId           string            `json:"tokenId"`
	Amount            string            `json:"amount"`
	AvailableAmount   string            `json:"availableAmount"`
	Attributes        map[string]string `json:"attributes"`
	Description       string            `json:"description"`
	Image             string            `json:"image"`
	Name              string            `json:"name"` // NFT名+id
	Platform          string            `json:"platform"`
} // 欧意订单模型

type ProtocolDataModel struct {
	Parameters ProtocolDataParametersModel `json:"parameters"`
	Signature  string                      `json:"signature"`

	// case2
	BasicColections []struct {
		Items []struct {
			Erc20TokenAmount string `json:"erc20TokenAmount"`
			NFTID            string `json:"nftId"`
		}
		NFTAddress          string `json:"nftAddress"`
		PlatFormFee         string `json:"platFormFee"`
		RoyaltyFee          int    `json:"royaltyFee"`
		RoyaltyFeeRecipient string `json:"royaltyFeeRecipient"`
	}
	Collections          interface{} `json:"collections"` // null 值使用 interface{} 处理
	ExpirationTime       int64       `json:"expirationTime"`
	ExtraData            string      `json:"extraData"`
	Hash                 string      `json:"hash"`
	HashNonce            string      `json:"hashNonce"`
	ListingTime          int64       `json:"listingTime"`
	Maker                string      `json:"maker"`
	Nonce                int         `json:"nonce"`
	OracleSignature      int         `json:"oracleSignature"`
	PaymentToken         string      `json:"paymentToken"`
	PlatformFeeRecipient string      `json:"platformFeeRecipient"`
	R                    string      `json:"r"`
	S                    string      `json:"s"`
	SignatureType        int         `json:"signatureType"`
	StartNonce           int         `json:"startNonce"`
	V                    int         `json:"v"`
}

type ProtocolDataParametersModel struct {
	ConduitKey                      string               `json:"conduitKey"`
	Consideration                   []ConsiderationModel `json:"consideration"` // 出价单
	Counter                         string               `json:"counter"`
	EndTime                         string               `json:"endTime"`
	Offer                           []OfferModel         `json:"offer"` // 挂单
	Offerer                         string               `json:"offerer"`
	OrderType                       OrderType            `json:"orderType"`
	Salt                            string               `json:"salt"`
	StartTime                       int64                `json:"startTime"`
	TotalOriginalConsiderationItems int                  `json:"totalOriginalConsiderationItems"`
	Zone                            string               `json:"zone"`
	ZoneHash                        string               `json:"zoneHash"`
}

type OrderParamsModel struct {
	Offerer                         string               `json:"offerer"`       // 订单发起人的地址
	Offer                           []OfferModel         `json:"offer"`         // 挂单项目模型数组，具体结构待定义
	Consideration                   []ConsiderationModel `json:"consideration"` // 订单项目模型数组，具体结构待定义
	StartTime                       int64                `json:"startTime"`
	EndTime                         int64                `json:"endTime"`
	OrderType                       OrderType            `json:"orderType"`
	Zone                            string               `json:"zone"` // 0x0000000000000000000000000000000000000000000000000000000000000000
	ZoneHash                        string               `json:"zoneHash"`
	Salt                            string               `json:"salt"`                            // 随机数
	ConduitKey                      string               `json:"conduitKey"`                      // 0x066003C1493A346357Af15158cD985b4A6e29D3F888888888888888888888888
	TotalOriginalConsiderationItems int                  `json:"totalOriginalConsiderationItems"` // 订单列表数量
	Counter                         int                  `json:"counter"`                         // 订单发起次数的计数器
} // 订单参数模型

type OfferModel struct {
	ItemType             ItemType `json:"itemType"`
	Token                string   `json:"token"`
	IdentifierOrCriteria string   `json:"identifierOrCriteria"`
	StartAmount          string   `json:"startAmount"`
	EndAmount            string   `json:"endAmount"`
} // 挂单项目类型

type ConsiderationModel struct {
	ItemType             ItemType `json:"itemType"`
	Token                string   `json:"token"`
	IdentifierOrCriteria string   `json:"identifierOrCriteria"`
	StartAmount          string   `json:"startAmount"`
	EndAmount            string   `json:"endAmount"`
	Recipient            string   `json:"recipient"`
} // 出价单项目类型

// OrderType 枚举类型
type OrderType int

// 枚举值
const (
	FULL_OPEN OrderType = iota
	PARTIAL_OPEN
	FULL_RESTRICTED
	PARTIAL_RESTRICTED
)

// ItemType 枚举类型
type ItemType int

// 枚举值
const (
	NATIVE ItemType = iota
	ERC20
	ERC721
	ERC1155
	ERC721_WITH_CRITERIA
	ERC1155_WITH_CRITERIA
)

// OrderType
// FULL_OPEN, PARTIAL_OPEN, FULL_RESTRICTED, PARTIAL_RESTRICTED
// ItemType
// NATIVE, ERC20, ERC721, ERC1155, ERC721_WITH_CRITERIA, ERC1155_WITH_CRITERIA

// CreateOffer创建订单
type OrderModel struct {
	CollectionAddress string      `json:"collectionAddress"`
	Count             string      `json:"count"`
	CurrencyAddress   string      `json:"currencyAddress"`
	ID                string      `json:"id"`
	ListingProfit     string      `json:"listingProfit"`
	NFTID             string      `json:"nftId"`
	Platform          string      `json:"platform"`
	PlatformFeePoints interface{} `json:"platformFeePoints"`
	Price             string      `json:"price"`
	Project           string      `json:"project"`
	ProtocolFeePoints interface{} `json:"protocolFeePoints"`
	RoyaltyFeePoints  interface{} `json:"royaltyFeePoints"`
	Source            int         `json:"source"`
	TokenID           string      `json:"tokenId"`
	ValidTime         int64       `json:"validTime"`
}

type OrderStepModel struct {
	Action string `json:"action"`
	Items  []struct {
		Amount            string   `json:"amount"`
		ApprovalAddress   string   `json:"approvalAddress"`
		Chain             int      `json:"chain"`
		Currency          string   `json:"currency"`
		CurrencyUrl       string   `json:"currencyUrl"`
		ContractAddress   string   `json:"contractAddress"`
		CollectionAddress string   `json:"collectionAddress"`
		Description       string   `json:"description"`
		Input             string   `json:"input"`
		Kind              string   `json:"kind"`
		OrderIDs          []string `json:"orderIds"`
		Platform          struct {
			Icon   string `json:"icon"`
			Name   string `json:"name"`
			Source string `json:"source"`
		} `json:"platform"`
		Platforms []struct {
			Icon   string `json:"icon"`
			Name   string `json:"name"`
			Source string `json:"source"`
		} `json:"platforms"`
		Status        string `json:"status"`
		TokenAddress  string `json:"tokenAddress"`
		TotalPrice    string `json:"totalPrice"`
		TotalUsdPrice string `json:"totalUsdPrice"`
		Value         string `json:"value"`

		Data   ProtocolDataParametersModel `json:"data"`
		Domain struct {
			ChainID           int    `json:"chainId"`
			Name              string `json:"name"`
			VerifyingContract string `json:"verifyingContract"`
			Version           string `json:"version"`
		} `json:"domain"`
		Post        interface{} `json:"post"`
		PrimaryType string      `json:"primaryType"`
		SignKind    string      `json:"signKind"`
		Types       Types       `json:"types"`
	} `json:"items"`
}

type Types struct {
	ConsiderationItem []struct {
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"ConsiderationItem"`
	OrderComponents []struct {
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"OrderComponents"`
	EIP712Domain []struct {
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"EIP712Domain"`
	OfferItem []struct {
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"OfferItem"`
}
