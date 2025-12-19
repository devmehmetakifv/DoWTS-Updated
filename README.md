# DoWTS
Denial of Wallet Test Simulator

**Last Updated: December 2025**

## Overview

DoWTS is a simulation tool for testing "Denial of Wallet" attack scenarios on serverless computing platforms. It calculates theoretical costs based on official pricing models without requiring actual cloud deployments.

## Features

- **Multi-Cloud Support**: Simulates attacks on AWS Lambda, Azure Functions, Google Cloud Functions, and IBM Cloud Functions
- **Local Simulation**: No cloud deployment required - all calculations run locally
- **Usage Pattern Generation**: 
  - Synthetic usage generation
  - Dataset-based patterns
  - Attack scenario simulation
- **Cost Analysis**: Real-time cost tracking based on current (Dec 2025) cloud provider pricing

## Supported Platforms & Pricing (Updated Dec 2025)

| Platform | Compute Price ($/GB-sec) | Request Price ($/million) | Free Tier |
|----------|--------------------------|---------------------------|-----------|
| AWS Lambda | $0.0000166667 | $0.20 | 1M requests, 400K GB-s |
| Azure Functions | $0.000016 | $0.20 | 1M requests, 400K GB-s |
| Google Cloud Functions (1st gen) | $0.0000025 | $0.40 | 2M requests, 400K GB-s |
| IBM Cloud Functions* | $0.000017 | Free | 400K GB-s |

*Note: IBM Cloud Functions was deprecated in 2024. Pricing model retained for historical comparison.

## Requirements

- Go 1.23 or higher
- MongoDB (for logging results)

## Installation

```bash
git clone https://github.com/yourusername/DoWTS
cd DoWTS
go mod download
```

## Usage

### Auto Mode (Randomized Parameters)
```bash
go run dowts.go 1
```

### Manual Mode (Custom Parameters)
```bash
go run dowts.go <input1> <input2> <input3> <input4> <input5> <input6>
```

## Architecture

```
DoWTS/
├── dowts.go              # Main entry point
├── go-packages/
│   ├── AWSLambda/        # AWS Lambda pricing simulation
│   ├── AzureFunctions/   # Azure Functions pricing simulation
│   ├── GoogleFunctions/  # Google Cloud Functions pricing
│   ├── IBMFunctions/     # IBM Cloud Functions pricing
│   ├── function/         # Generic function model
│   └── UsageGenerator/   # Usage pattern generators
│       ├── AttackUsageGenerator.go
│       ├── DatasetUsageGenerator.go
│       ├── SyntheticUsageGenerator.go
│       └── UsageGenerator.go
```

## Important Notes

- **Simulation Only**: This tool does NOT deploy actual functions to cloud platforms
- **Pricing Accuracy**: Pricing updated December 2025. Always verify with official cloud provider documentation
- **Research Tool**: Designed for academic research and security analysis
- **No Cloud Costs**: Running this tool incurs no cloud computing charges

## Recent Updates (Dec 2025)

- ✅ Updated to Go 1.23
- ✅ Verified AWS Lambda pricing ($0.0000166667/GB-s)
- ✅ Verified Azure Functions pricing ($0.000016/GB-s)
- ✅ Verified Google Cloud Functions 1st gen pricing ($0.0000025/GB-s)
- ✅ Added deprecation notice for IBM Cloud Functions
- ⚠️ Note: Google Cloud Functions 2nd gen (Cloud Run-based) has different pricing model

## License

[Add your license here]

## Citation

If you use this tool in your research, please cite appropriately.
