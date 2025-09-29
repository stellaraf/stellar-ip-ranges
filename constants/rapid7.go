package constants

import "go.stellar.af/stellar-ip-ranges/types"

// Rapid7 networks/domains formatted for types.List
var R7_IP4 = types.List{
    // Agent / Ingress endpoints (US region)
    "34.226.68.35",
    "54.144.111.231",
    "52.203.25.223",
    "34.236.161.191",

    // Main ingress subnet
    "193.149.136.0/24",

    // CDN / Content delivery IPs
    "3.163.232.9",
    "3.163.233.9",
    "3.163.234.9",
    "3.163.235.9",
    "3.163.236.9",
    "3.163.237.9",
    "3.163.238.9",
    "3.163.239.9",
    "3.163.240.9",
    "3.163.241.9",
    "3.163.242.9",
    "3.163.243.9",
    "3.163.244.9",
    "3.163.245.9",
    "3.163.246.9",
    "3.163.247.9",
    "3.163.248.9",
    "3.163.249.9",
    "3.163.250.9",
    "3.163.251.9",
    "3.163.252.9",
}

// IPv6 addresses are not published by Rapid7
var R7_IP6 = types.List{
}

var R7_DOMAINS = types.List{
    // Core Insight Platform
    "*.insight.rapid7.com",
    "data.insight.rapid7.com",
    "data.logs.insight.rapid7.com",

    // Agent / Ingress endpoints
    "*.endpoint.ingress.rapid7.com",
    "endpoint.ingress.rapid7.com",
    "us.endpoint.ingress.rapid7.com",
    "us.api.endpoint.ingress.rapid7.com",
    "us.bootstrap.endpoint.ingress.rapid7.com",
    "us.deployment.endpoint.ingress.rapid7.com",
    "us.main.endpoint.ingress.rapid7.com",
    "us.storage.endpoint.ingress.rapid7.com",
    "us.storage.main.endpoint.ingress.rapid7.com",
    "us.api.main.endpoint.ingress.rapid7.com",

    // Regional data endpoints
    "us2.data.insight.rapid7.com",
    "us3.data.insight.rapid7.com",
    "ca.data.insight.rapid7.com",
    "eu.data.insight.rapid7.com",
    "ap.data.insight.rapid7.com",
    "au.data.insight.rapid7.com",

    // Regional deployment endpoints
    "us2.deployment.endpoint.ingress.rapid7.com",
    "us3.deployment.endpoint.ingress.rapid7.com",
    "ca.deployment.endpoint.ingress.rapid7.com",
    "eu.deployment.endpoint.ingress.rapid7.com",
    "ap.deployment.endpoint.ingress.rapid7.com",
    "au.deployment.endpoint.ingress.rapid7.com",

    // Amazon S3 endpoints used by collectors
    "s3.amazonaws.com",
    "s3.us-east-2.amazonaws.com",
    "s3.us-west-2.amazonaws.com",
    "s3.ca-central-1.amazonaws.com",
    "s3.eu-central-1.amazonaws.com",
    "s3-ap-northeast-1.amazonaws.com",
    "s3-ap-southeast-2.amazonaws.com",
}
