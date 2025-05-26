# Cloud Run Simple

This example demonstrates how to define and use **Pigen plugins** and **steps** to deploy a containerized service to **Google Cloud Run**, leveraging **Artifact Registry** and **Secret Manager**.

---

## 📦 Pigen Plugins File

```yaml
plugins:

  # Artifact Registry Plugin
  # Artifact Registry is a fully managed service for storing container images and other build artifacts.
  - id: artifact_registry_tf
    repo_url: https://github.com/pigen-dev/artifact-registry-tf-plugin
    version: v0.1.2
    plugin:
      label: ARTIFACT_REGISTRY_DEMO
      config:
        location:         # Set your region
        repo_id:          # Set your repository name
        project_id:       # Set your project ID
        description:      # (Optional) Description of the repository
      output:
        repo_url:

  # Google Cloud Run Plugin
  # Cloud Run is a fully managed platform for deploying and scaling containerized applications.
  - id: google_cloud_run
    repo_url: https://github.com/pigen-plugins/google-cloud-run
    version: v0.2.1
    plugin:
      label: GOOGLE_CLOUD_RUN_DEMO
      config:
        location:         # Set your region
        project_id:       # Set your project ID
        service_name:     # Set your service name
        unauthenticated: true
      output:
        cloud_run_url:
        service_name:

  # Secret Manager Plugin
  # Secret Manager stores sensitive data like API keys and credentials securely.
  - id: secret_manager
    repo_url: https://github.com/pigen-plugins/secret-manager
    version: v0.1.3
    plugin:
      label: SECRET_MANAGER_DEMO
      config:
        project_id:       # Set your project ID
        prefix:           # Prefix used to format secret names
        secrets:
          # Values are loaded from a `.env.pigen` file
          GOOGLE_API_KEY: {{.ENV.GOOGLE_API_KEY}}
      output:
        secrets_list: []
        prefix:
```

### ✅ This configuration sets up the following infrastructure:
- **Artifact Registry**: Creates a Docker repository to store container images.
- **Google Cloud Run**: Deploys your container to a serverless managed service.
- **Secret Manager**: Stores secrets securely. Use a `.env.pigen` file for secret values, and add it to `.gitignore` to keep secrets private.

---

## ⚙️ Pigen Steps File

```yaml
type: cloudbuild
version: v1.1.1
repo_url: https://github.com/pigen-dev/cloudbuild-plugin
config:
  deployment:
    # NOTE: Use string format for project_number ("775465758731" not 775465758731)
    project_number: "775465758731"
    project_id: aidodev
    project_region: "europe-west1"
  github_url: https://github.com/pigen-dev/pigen-core.git
  target_branch: "^test-pigen$"

steps:
  # Step 1: Build and push Docker image to Artifact Registry
  - step: docker-build-push
    placeholders:
      image: "{{.Plugins.ARTIFACT_REGISTRY_DEMO.repo_url}}/landing:latest"

  # Step 2: Deploy the container to Google Cloud Run
  - step: google-cloud-run
    placeholders:
      service_name: {{.Plugins.GOOGLE_CLOUD_RUN_DEMO.service_name}}
      image: "{{.Plugins.ARTIFACT_REGISTRY_DEMO.repo_url}}/landing:latest"
      secrets:
        secret_list: {{.Plugins.SECRET_MANAGER_DEMO.secrets_list}}
        secret_prefix: {{.Plugins.SECRET_MANAGER_DEMO.prefix}}
```

---

## 🚀 Pipeline Overview

This pipeline uses **Cloud Build** as the CI/CD tool and includes two main steps:

### 1. **docker-build-push**
Builds your Docker image and pushes it to the configured Artifact Registry repository.

### 2. **google-cloud-run**
Deploys the image to a Cloud Run service. It also injects environment secrets fetched from Secret Manager.

---

## 📝 Notes
- `github_url`: URL of your GitHub repository.
- `target_branch`: Used to trigger the pipeline on specific branches (e.g., `test-pigen`).
- `.env.pigen`: File for secret values. **Never commit this file**; add it to your `.gitignore`.