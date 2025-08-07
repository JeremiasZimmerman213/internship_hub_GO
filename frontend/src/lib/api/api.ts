const API_BASE_URL = 'http://localhost:8080';

export async function createApplication(formData = new FormData()) {
  const response = await fetch(`${API_BASE_URL}/applications`, {
    method: 'POST',
    body: formData, // FormData automatically sets correct Content-Type
  });

  if (!response.ok) {
    const errorData = await response.json().catch(() => ({ error: 'Unknown error occurred' }));
    throw new Error(errorData.error || `HTTP ${response.status}: ${response.statusText}`);
  }

  return await response.json();
}

export async function getApplications() {
  const response = await fetch(`${API_BASE_URL}/applications`);
  
  if (!response.ok) {
    const errorData = await response.json().catch(() => ({ error: 'Unknown error occurred' }));
    throw new Error(errorData.error || `HTTP ${response.status}: ${response.statusText}`);
  }

  return await response.json();
}

export async function getApplication(id: string | number) {
  const response = await fetch(`${API_BASE_URL}/applications/${id}`);
  
  if (!response.ok) {
    const errorData = await response.json().catch(() => ({ error: 'Unknown error occurred' }));
    throw new Error(errorData.error || `HTTP ${response.status}: ${response.statusText}`);
  }

  return await response.json();
}

export async function deleteApplication(id: string | number) {
  const response = await fetch(`${API_BASE_URL}/applications/${id}`, {
    method: 'DELETE',
  });

  if (!response.ok) {
    const errorData = await response.json().catch(() => ({ error: 'Unknown error occurred' }));
    throw new Error(errorData.error || `HTTP ${response.status}: ${response.statusText}`);
  }

  // DELETE responses often return empty body, so check if there's content
  const contentType = response.headers.get('content-type');
  if (contentType && contentType.includes('application/json')) {
    return await response.json();
  }
  return { success: true };
}


/**
 * @param {string | number} id
 * @param {FormData} formData
 */
export async function updateApplication(id: string | number, formData: FormData) {
    const response = await fetch(`${API_BASE_URL}/applications/${id}`, {
        method: 'PUT',
        body: formData,
    });

    if (!response.ok) {
        const errorData = await response.json().catch(() => ({ error: 'Unknown error occurred' }));
        throw new Error(errorData.error || `HTTP ${response.status}: ${response.statusText}`);
    }

    return await response.json();
}
