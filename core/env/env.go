package env

var JwtSecretKet []byte

var (
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
)

var HttpPort string

var (
	AwsProfileImageBucketName      string
	AwsProfileImageAccessKey       string
	AwsProfileImageSecretAccessKey string
	AwsProfileImageFolderPath      string
	AwsProfileImageRegion          string
)
