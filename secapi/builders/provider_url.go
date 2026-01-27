package builders

type ProviderUrlBuilder struct {
	baseDomain     string
	pathProvider   string
	domainProvider string
	version        string
}

func NewProviderUrlBuilder() *ProviderUrlBuilder {
	return &ProviderUrlBuilder{}
}

func (b *ProviderUrlBuilder) BaseDomain(baseDomain string) *ProviderUrlBuilder {
	b.baseDomain = baseDomain
	return b
}

func (b *ProviderUrlBuilder) PathProvider(pathProvider string) *ProviderUrlBuilder {
	b.pathProvider = pathProvider
	return b
}

func (b *ProviderUrlBuilder) DomainProvider(domainProvider string) *ProviderUrlBuilder {
	b.domainProvider = domainProvider
	return b
}

func (b *ProviderUrlBuilder) Version(version string) *ProviderUrlBuilder {
	b.version = version
	return b
}

func (b *ProviderUrlBuilder) BuildPathUrl() string {
	return "https://" + b.baseDomain + "/providers/" + b.pathProvider + "/" + b.version + "/"
}

func (b *ProviderUrlBuilder) BuildDomainUrl() string {
	return "https://" + b.domainProvider + "." + b.baseDomain + "/" + b.version + "/"
}
