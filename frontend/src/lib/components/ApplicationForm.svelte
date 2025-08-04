<script lang="ts">
  import { goto } from '$app/navigation';
  
  export let onSubmit: (formData: FormData) => Promise<void> = async () => {};
  export let isLoading = false;
  
  let company = '';
  let position = '';
  let status = 0; // Default to "Applied"
  let location = '';
  let appliedDate = '';
  let term = '';
  let note = '';
  let resumeFile: File | null = null;
  let error = '';
  let fileInputElement: HTMLInputElement;

  const statusOptions = [
    { value: 0, label: 'Applied' },
    { value: 1, label: 'Interview' },
    { value: 2, label: 'Offer Received' },
    { value: 3, label: 'Rejected' }
  ];

  function validateForm(): boolean {
    error = '';
    
    if (!company.trim()) {
      error = 'Company name is required.';
      return false;
    }
    if (!position.trim()) {
      error = 'Position is required.';
      return false;
    }
    if (!location.trim()) {
      error = 'Location is required.';
      return false;
    }
    if (!appliedDate) {
      error = 'Applied date is required.';
      return false;
    }
    if (!term.trim()) {
      error = 'Term is required.';
      return false;
    }
    if (!resumeFile) {
      error = 'Resume PDF is required.';
      return false;
    }
    if (resumeFile && resumeFile.type !== 'application/pdf') {
      error = 'Resume must be a PDF file.';
      return false;
    }
    if (resumeFile && resumeFile.size > 5 * 1024 * 1024) {
      error = 'Resume file must be less than 5MB.';
      return false;
    }
    
    return true;
  }

  function handleFileChange(event: Event) {
    const target = event.target as HTMLInputElement;
    resumeFile = target.files?.[0] || null;
  }

  async function handleSubmit(event: Event) {
    event.preventDefault();
    
    if (!validateForm()) return;
    
    const formData = new FormData();
    formData.append('company', company.trim());
    formData.append('position', position.trim());
    formData.append('status', status.toString());
    formData.append('location', location.trim());
    formData.append('applied_date', new Date(appliedDate).toISOString());
    formData.append('term', term.trim());
    formData.append('note', note.trim());
    if (resumeFile) {
      formData.append('resume', resumeFile);
    }
    
    try {
      await onSubmit(formData);
    } catch (err) {
      error = err instanceof Error ? err.message : 'An error occurred while submitting the form.';
    }
  }

  function resetForm() {
    company = '';
    position = '';
    status = 0;
    location = '';
    appliedDate = '';
    term = '';
    note = '';
    resumeFile = null;
    error = '';
    if (fileInputElement) {
      fileInputElement.value = '';
    }
  }
</script>

<form on:submit|preventDefault={handleSubmit} class="application-form">
  <div class="row">
    <div class="col-md-6 mb-3">
      <label for="company" class="form-label">Company Name *</label>
      <input 
        type="text" 
        class="form-control" 
        id="company" 
        bind:value={company}
        required 
        disabled={isLoading}
        placeholder="e.g. Google, Microsoft"
      />
    </div>
    <div class="col-md-6 mb-3">
      <label for="position" class="form-label">Position *</label>
      <input 
        type="text" 
        class="form-control" 
        id="position" 
        bind:value={position}
        required 
        disabled={isLoading}
        placeholder="e.g. Software Engineer Intern"
      />
    </div>
  </div>

  <div class="row">
    <div class="col-md-6 mb-3">
      <label for="status" class="form-label">Status *</label>
      <select 
        class="form-select" 
        id="status" 
        bind:value={status}
        required 
        disabled={isLoading}
      >
        {#each statusOptions as option}
          <option value={option.value}>{option.label}</option>
        {/each}
      </select>
    </div>
    <div class="col-md-6 mb-3">
      <label for="location" class="form-label">Location *</label>
      <input 
        type="text" 
        class="form-control" 
        id="location" 
        bind:value={location}
        required 
        disabled={isLoading}
        placeholder="e.g. San Francisco, CA"
      />
    </div>
  </div>

  <div class="row">
    <div class="col-md-6 mb-3">
      <label for="appliedDate" class="form-label">Applied Date *</label>
      <input 
        type="date" 
        class="form-control" 
        id="appliedDate" 
        bind:value={appliedDate}
        required 
        disabled={isLoading}
      />
    </div>
    <div class="col-md-6 mb-3">
      <label for="term" class="form-label">Term *</label>
      <input 
        type="text" 
        class="form-control" 
        id="term" 
        bind:value={term}
        required 
        disabled={isLoading}
        placeholder="e.g. Summer 2025, Fall 2024"
      />
    </div>
  </div>

  <div class="mb-3">
    <label for="note" class="form-label">Notes</label>
    <textarea 
      class="form-control" 
      id="note" 
      bind:value={note}
      rows="3"
      disabled={isLoading}
      placeholder="Any additional notes about this application..."
      maxlength="1048"
    ></textarea>
    <div class="form-text">{note.length}/1048 characters</div>
  </div>

  <div class="mb-3">
    <label for="resume" class="form-label">Resume (PDF) *</label>
    <input 
      type="file" 
      class="form-control" 
      id="resume" 
      accept=".pdf"
      required 
      disabled={isLoading}
      bind:this={fileInputElement}
      on:change={handleFileChange}
    />
    <div class="form-text">Max file size: 5MB. Only PDF files are allowed.</div>
    {#if resumeFile}
      <div class="mt-2">
        <small class="text-success">
          ✓ Selected: {resumeFile.name} ({(resumeFile.size / 1024 / 1024).toFixed(2)} MB)
        </small>
      </div>
    {/if}
  </div>

  {#if error}
    <div class="alert alert-danger py-2 mb-3">{error}</div>
  {/if}

  <div class="d-flex gap-2">
    <button 
      type="submit" 
      class="btn flex-fill"
      style="background-color: #B1B2FF; color: #2d2d2d; font-weight: 600;"
      disabled={isLoading}
    >
      {#if isLoading}
        <span class="spinner-border spinner-border-sm me-2" role="status"></span>
        Creating...
      {:else}
        Create Application
      {/if}
    </button>
    <button 
      type="button" 
      class="btn btn-outline-secondary"
      on:click={resetForm}
      disabled={isLoading}
    >
      Reset
    </button>
  </div>
</form>

<style>
.application-form {
  background: #EEF1FF;
  padding: 2rem;
  border-radius: 1rem;
  border: none;
}

.form-control:focus, .form-select:focus {
  border-color: #B1B2FF;
  box-shadow: 0 0 0 0.2rem rgba(177, 178, 255, 0.25);
}

.btn:hover {
  background-color: #AAC4FF !important;
}
</style>
