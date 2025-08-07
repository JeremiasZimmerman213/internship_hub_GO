const API_BASE_URL = 'http://localhost:8080';

/**
 * @param {FormData} formData
 */
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

/**
 * @param {string | number} id
 */
export async function getApplication(id) {
  const response = await fetch(`${API_BASE_URL}/applications/${id}`);
  
  if (!response.ok) {
    const errorData = await response.json().catch(() => ({ error: 'Unknown error occurred' }));
    throw new Error(errorData.error || `HTTP ${response.status}: ${response.statusText}`);
  }

  return await response.json();
}

/**
 * @param {string | number} id
 */
export async function deleteApplication(id) {
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
