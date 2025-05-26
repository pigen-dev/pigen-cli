# PIGEN CLI

The **PIGEN CLI** is a lightweight, command-line tool that interfaces directly with **PIGEN Core**, enabling users to generate CI/CD pipelines and manage plugins through a local terminal. It is intended for users who prefer or require scriptable, UI-free interaction with PIGEN.

The CLI version offers a **free and open-source** alternative to the full PIGEN platform. However, it does **not include access to the web dashboard or AI-powered features** (such as Pi-Pilot). These advanced capabilities are exclusive to the **enterprise monetized version**, which provides a more complete DevOps experience through PIGEN Web.

---

## 🚀 Use Cases

- **Fast Pipeline Scaffolding**: Bootstrap CI/CD pipelines in seconds using predefined or custom steps.
- **Offline Usage**: Generate and modify pipelines without needing internet access or a web interface.
- **Scripting and Automation**: Integrate PIGEN into custom scripts or DevOps automation workflows.
- **Plugin Management**: Install, update, and configure plugins that enhance pipeline steps or cloud integrations.
- **Version Control Friendly**: All pipeline definitions are saved as files, ideal for Git-based workflows.

---

# ⚙️ How to Use

## Installation

#### Linux/macOS Installation

1. **Download the Precompiled Binary:**

   Visit the [Releases](https://github.com/pigen-dev/pigen-cli/releases) page and download the appropriate binary for your operating system.

2. **Install the Binary:**

   Move the binary to the `/usr/local/bin` directory to make it globally accessible from your terminal:

   ```bash
   # Make sure the binary has execute permissions
   chmod +x pigen

   # Move it to /usr/local/bin
   sudo mv pigen /usr/local/bin/

3. **Prepare your Environment:**

    The cli doesn't take actions on your cloud environment directly it uses the pigen-core that should be deployed in your cloud environment. this is done by running the pigen command:
     ```bash
     # Run the following command and fill in the form
     pigen config init
This will deploy the pigen-core instance in your cloud (for GCP it's deployed on cloud run). The instance endpoint and details will be stored in ~/.pigen file.

---

## 🔌 PiGen Plugins

Welcome to the **PiGen Plugin System**!  
Plugins are the core mechanism that allow PiGen to extend, customize, and automate DevOps and cloud infrastructure in a simple, modular way.

---

### 📦 What Is a Plugin?

A **plugin** in PiGen represents a **cloud resource**, **DevOps tool**, or **infrastructure component**. It abstracts away complex configuration logic (like Terraform, Kubernetes manifests, or Bash scripts) into a reusable and customizable unit.

With plugins, you can define what you need in a clean YAML format without worrying about implementation details.

---

### ✅ What Can Plugins Do?

PiGen plugins can:

- Provision cloud resources  
  _e.g., GKE clusters, storage buckets, artifact registries_
- Deploy DevOps tools  
  _e.g., ArgoCD, Helm, Prometheus_
- Automate DevOps operations  
  _e.g., sending HTTP notifications, setting up CI/CD steps_
- Integrate third-party services and tools

---

### ⚙️ How Do Plugins Work?

- Each plugin defines a **schema** (inputs and outputs).
- Users configure plugins using YAML and pass required parameters.
- Plugins are executed locally or remotely (using [HashiCorp Go-Plugin](https://github.com/hashicorp/go-plugin)).
- They generate code, infrastructure, or actions that integrate directly into PiGen pipelines.

---

### 🧰 Plugin Examples

| Plugin Name        | Description                                         |
|--------------------|-----------------------------------------------------|
| `gke-cluster`      | Provisions a Google Kubernetes Engine cluster       |
| `argocd`           | Installs and configures ArgoCD in a Kubernetes env  |
| `artifact-registry`| Creates a secure artifact registry on GCP           |
| `http-notifier`    | Sends a webhook notification to a given URL         |

---

## 🚀 Installing a Plugin

### 1. Browse the Plugin Hub

Visit the **[PiGen Hub](#)** to explore a curated list of plugins, categorized into:

- 🛠️ **DevOps**
- 💾 **Data & Storage**
- 🤖 **MLOps**
- 🔐 **Security & Secrets**
- And more...

Each plugin in the hub includes a description, YAML configuration snippet, and a breakdown of required inputs.

---

### 2. Add the Plugin to Your `pigen-plugins.yaml`

Here is an example of `pigen-plugins.yaml` file:

```yaml
plugins:
- id: artifact_registry_tf
  repo_url: https://github.com/pigen-dev/artifact-registry-tf-plugin
  version: v0.1.1
  plugin:
    label: ARTIFACT_REGISTRY_DEMO
    config:
      location: # Your cloud location where to provision this ressource
      repo_id: # Repo name
      project_id: # Your cloud project ID
      description: # Set a description for your repo (optional)
    output:
      repo_url: # This an output only leave it blank. Your created repo endpoint URL
```
#### ⚠️ Note:
- Paste the plugin YAML config under the plugins: key.

- Replace input placeholders with your values.

- Each plugin must have a **unique label** to avoid conflicts.

### 3.  Install your Plugins
   Once you've added all your desired plugins:
   ```bash
   pigen plugin install
   ```
   This command will install all configured plugins. Already-installed plugins will be skipped to prevent duplication.

### 4.  Get Plugin Output
   To retrieve output values from a specific plugin (e.g., an IP address, bucket URL, etc.):
   ```bash
   pigen plugin output {plugin-label}
   ```

### 5.  Destroy a Plugin
   To remove and clean up resources associated with a plugin:
   ```bash
   pigen plugin destroy {plugin-label}
   ```
   
---


## 🪜 PiGen Steps
PiGen Steps are reusable, declarative actions that define the logic of your CI/CD pipelines. While Plugins focus on provisioning resources or tools, Steps focus on how your pipeline will behave.

## 🧩 What Is a PiGen Step?
A PiGen Step is a template that defines a single pipeline step (e.g., build, test, deploy).
It is:

- **Tool-agnostic:** One step definition can support multiple CI/CD tools.

- **Templated:** Contains placeholders (e.g., `image`, `secrets`) that are dynamically filled during pipeline generation.

- **Reusable:** Can be shared cross different projects.
---
### Pigen Step structure:
Each Pigen Step follows this structure:
```yaml
  - step: #step id
    placeholders:
      #list of placeholders that will parsed in the template
```
This is a simple example:
```yaml
- step: docker-build-push
  placeholders:
    image: pigen/testing:latest
```
---
### 📂Pigen Step Hub:
Similar to Pigen Plugins Hub, Pigen Steps Hub represents a list of pigen steps that you can use to build your reusable pipelines.
Each step in the hub includes a description, YAML configuration snippet, and a breakdown of required placeholders.
## 🛠 How to Setup Your Pipeline

### Pigen steps file
Similar to Pigen Plugins to configure your Pipeline you simple create a `pigen-steps.yaml` file where you define:
- Your CICD tool
- Your github repo URL
- Your target branch
- Your CICD steps
---
## 🔧 Configure Your Pipeline:
Pigen supports multiple cicd tools. For now we only support Google CloudBuild.
In your `pigen-steps.yaml` file:

```yaml
type: cloudbuild
version: v1.1.0
repo_url: https://github.com/pigen-dev/cloudbuild-plugin
config:
  deployment:
    project_number: "775465758731"
    project_id: aidodev
    project_region: "europe-west1"
  github_url: https://github.com/pigen-dev/pigen-core.git
  target_branch: "^test-pigen$"
steps:
  - step: docker-build-push
    placeholders:
      image: pigen/testing:latest
```
- **type** is your desired CICD tool (`cloudbuild`, `github-actions`, `circle_ci`...)
- **version** is the CICD tool plugin release version
- **repo_url** is the CICD tool plugin github url
- **config** is the CICD tool configuration and this depends on the CICD plugin itself
- **steps** is the list of pigen steps that you want to use in your pipeline
#### ⚠️ Note:
- The CICD tool is a special plugin that will be installed in your cloud environment.
- The CICD tool block can be found also in [Plugin Hub](#) which makes it easy for you to write your `pigen-steps.yaml` file 
---
### 1.  Setup Your CICD tool
  Deploy your cicd tool plugin if it needs to and create a trigger on your branch
   ```bash
   pigen pipeline setup
   ```

### 2.  Generate Your Pipeline Script
   To finally generate your pipeline script:
   ```bash
   pigen pipeline generate
   ```

---
### 🔗 Use Your Plugin Output

Using Pigen you will be able to directly link your installed plugins to your pipeline.

Each insalled plugin exposes a list of outputs that you can get access to following this structure `{{.Plugins.plugin_label.output_key}}`.

The output value can be used in your pigen yaml files. For example `artifact_registry_tf` plugin exposes an output `repo_url` that can be used in your `docker-build-push` step to forme your `image` placeholder:

```yaml
steps:
  - step: docker-build-push
    placeholders:
      image: "{{.Plugins.ARTIFACT_REGISTRY_DEMO.repo_url}}/testing:latest"
```
Doing so the step image placeholder will take as value the **artifact registry** created **repo url** concatenated with `/testing:latest`

---
### 🔑 Use Secret Values in Pigen files

As Pigen files will probably get pushed in your git repo you don't want to share sensitive data like API keys or passwords.

That's when you want to use `.env.pigen` file to define all your sensitive data that you want to use to configure your `pigen-steps.yaml` or `pigen-plugins.yaml` files.

You can Simple define your keys in the `.env.pigen` file and reference it in your yaml file following this syntax `{{.ENV.your_key}}`

For example to configure `secret_manager` plugin you need to pass a list of key value pair of secrets that you don't want to be exposed in your yaml file so you simple define the values in the `.env.pigen` file:
```env
GOOGLE_API_KEY="testing it"
```
And you can use it in your `pigen-plugins.yaml` file:

```yaml
- id: secret_manager
  repo_url: https://github.com/pigen-plugins/secret-manager
  version: v0.1.3
  plugin:
    label: SECRET_MANAGER_DEMO
    config:
        project_id: aidodev
        prefix: PIGEN
        secrets:
          GOOGLE_API_KEY: {{.ENV.GOOGLE_API_KEY}}
    output:
      secrets_list: []
      prefix:
```

---

# Simple Example 
In [simples](https://github.com/pigen-dev/pigen-cli/simples) folder you can find some useful simples to test the tool.

[cloud-run-pipeline](https://github.com/pigen-dev/pigen-cli/simples/gcp/cloudrun-pipeline)  is a simple pipeline example to deploy a webapp on cloud run. Read the [README.md](#) file to better understand the example.


# 🌍 Community-Driven

PiGen plugins and steps are **open source** and **community-driven**.  
Contribute to the [Plugin Hub](#) and [Steps Hub](#), help shape the future of DevOps and MLOps automation!

---

# 🧑‍💻 Creating Your Own Plugin

Coming soon: A complete guide and template for creating and publishing your own plugins to the PiGen Plugin Hub.

---
# 📫 Feedback & Contribution

Found a bug or have an idea? Open an [issue](#) or contribute via [pull request](#).  
Together, let's make DevOps/MLOps simpler and smarter.

---

> © 2025 PiGen. Built with ♥ by pigen and the community.

   
