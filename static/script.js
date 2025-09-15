document.addEventListener('DOMContentLoaded', function() {
    // Initial page load - default to listing page
    loadMdsListingPage();
    
    // Event listener for MDS Entry link in sidebar
    document.getElementById('mdsEntryLink').addEventListener('click', function(e) {
        e.preventDefault();
        loadMdsEntryPage();
    });
});

function loadMdsEntryPage() {
    const contentArea = document.getElementById('content-area');
    
    contentArea.innerHTML = `
        <div class="page-header">
            <h1>MDS Entry Page</h1>
        </div>
        <div class="form-container">
            <form id="mdsEntryForm">
                <div class="form-field">
                    <label for="name" class="form-label required-field">Name</label>
                    <input type="text" class="form-control" id="name" required>
                    <div class="error-message" id="name-error"></div>
                </div>
                
                <div class="form-field">
                    <label for="comments" class="form-label">Comments</label>
                    <textarea class="form-control" id="comments" rows="3"></textarea>
                </div>
                
                <div class="form-field">
                    <label for="effectiveFrom" class="form-label required-field">Effective From</label>
                    <input type="date" class="form-control" id="effectiveFrom" required>
                    <div class="error-message" id="effectiveFrom-error"></div>
                </div>
                
                <div class="form-field">
                    <label for="effectiveTo" class="form-label required-field">Effective To</label>
                    <input type="date" class="form-control" id="effectiveTo" required>
                    <div class="error-message" id="effectiveTo-error"></div>
                </div>
                
                <div class="form-check form-field">
                    <input class="form-check-input" type="checkbox" id="isPPAgreed">
                    <label class="form-check-label" for="isPPAgreed">Is PP Agreed</label>
                </div>
                
                <div class="form-field">
                    <label for="documentPath" class="form-label">Document</label>
                    <input type="file" class="form-control" id="documentPath">
                </div>
                
                <div class="mt-4">
                    <button type="button" class="btn btn-primary btn-action" id="saveButton">Save</button>
                    <button type="button" class="btn btn-secondary btn-action" id="cancelButton">Cancel</button>
                </div>
            </form>
        </div>
    `;
    
    // Add event listeners for the form buttons
    document.getElementById('saveButton').addEventListener('click', saveMdsEntry);
    document.getElementById('cancelButton').addEventListener('click', function() {
        loadMdsListingPage();
    });
}

function loadMdsListingPage() {
    const contentArea = document.getElementById('content-area');
    
    contentArea.innerHTML = `
        <div class="page-header">
            <h1>MDS Entries</h1>
            <button class="btn btn-primary" id="addNewButton">Add New</button>
        </div>
        <div class="table-container">
            <table class="table table-striped">
                <thead>
                    <tr>
                        <th>ID</th>
                        <th>Name</th>
                        <th>Comments</th>
                        <th>Effective From</th>
                        <th>Effective To</th>
                        <th>PP Agreed</th>
                        <th>Actions</th>
                    </tr>
                </thead>
                <tbody id="mdsTableBody">
                    <tr>
                        <td colspan="7" class="text-center">Loading...</td>
                    </tr>
                </tbody>
            </table>
        </div>
    `;
    
    // Add event listener for Add New button
    document.getElementById('addNewButton').addEventListener('click', function() {
        loadMdsEntryPage();
    });
    
    // Load MDS entries
    fetchMdsEntries();
}

function saveMdsEntry() {
    // Clear previous error messages
    clearErrorMessages();
    
    // Get form values
    const name = document.getElementById('name').value.trim();
    const comments = document.getElementById('comments').value.trim();
    const effectiveFrom = document.getElementById('effectiveFrom').value;
    const effectiveTo = document.getElementById('effectiveTo').value;
    const isPPAgreed = document.getElementById('isPPAgreed').checked;
    const fileInput = document.getElementById('documentPath');
    let documentPath = '';
    
    if (fileInput.files.length > 0) {
        documentPath = fileInput.files[0].name; // In a real app, we would upload the file to a server
    }
    
    // Validation
    let isValid = true;
    
    if (!name) {
        document.getElementById('name-error').textContent = 'Name is required';
        isValid = false;
    }
    
    if (!effectiveFrom) {
        document.getElementById('effectiveFrom-error').textContent = 'Effective From date is required';
        isValid = false;
    }
    
    if (!effectiveTo) {
        document.getElementById('effectiveTo-error').textContent = 'Effective To date is required';
        isValid = false;
    }
    
    if (effectiveFrom && effectiveTo && new Date(effectiveTo) < new Date(effectiveFrom)) {
        document.getElementById('effectiveTo-error').textContent = 'Effective To date must not be earlier than Effective From date';
        isValid = false;
    }
    
    if (!isValid) {
        return;
    }
    
    // Prepare data for API
    const mdsEntry = {
        name: name,
        comments: comments,
        effectiveFrom: new Date(effectiveFrom).toISOString(),
        effectiveTo: new Date(effectiveTo).toISOString(),
        isPPAgreed: isPPAgreed,
        documentPath: documentPath
    };
    
    // Call API to save MDS entry
    fetch('/api/mds', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(mdsEntry)
    })
    .then(response => {
        if (!response.ok) {
            return response.json().then(data => {
                throw new Error(data.error || 'Failed to save MDS entry');
            });
        }
        return response.json();
    })
    .then(data => {
        alert('MDS entry has been saved successfully.');
        loadMdsListingPage();
    })
    .catch(error => {
        alert(error.message);
    });
}

function fetchMdsEntries() {
    fetch('/api/mds')
    .then(response => {
        if (!response.ok) {
            throw new Error('Failed to retrieve MDS entries');
        }
        return response.json();
    })
    .then(entries => {
        displayMdsEntries(entries);
    })
    .catch(error => {
        console.error('Error fetching MDS entries:', error);
        document.getElementById('mdsTableBody').innerHTML = `
            <tr>
                <td colspan="7" class="text-center">Error loading data: ${error.message}</td>
            </tr>
        `;
    });
}

function displayMdsEntries(entries) {
    const tableBody = document.getElementById('mdsTableBody');
    
    if (entries.length === 0) {
        tableBody.innerHTML = `
            <tr>
                <td colspan="7" class="text-center">No MDS entries found</td>
            </tr>
        `;
        return;
    }
    
    let tableContent = '';
    entries.forEach(entry => {
        const effectiveFrom = new Date(entry.effectiveFrom).toLocaleDateString();
        const effectiveTo = new Date(entry.effectiveTo).toLocaleDateString();
        
        tableContent += `
            <tr>
                <td>${entry.id}</td>
                <td>${entry.name}</td>
                <td>${entry.comments || '-'}</td>
                <td>${effectiveFrom}</td>
                <td>${effectiveTo}</td>
                <td>${entry.isPPAgreed ? 'Yes' : 'No'}</td>
                <td>
                    <button class="btn btn-sm btn-danger" onclick="deleteMdsEntry(${entry.id})">Delete</button>
                </td>
            </tr>
        `;
    });
    
    tableBody.innerHTML = tableContent;
}

function deleteMdsEntry(id) {
    if (confirm('Are you sure you want to delete this MDS entry?')) {
        fetch(`/api/mds/${id}`, {
            method: 'DELETE'
        })
        .then(response => {
            if (!response.ok) {
                throw new Error('Failed to delete MDS entry');
            }
            // Refresh the listing
            fetchMdsEntries();
        })
        .catch(error => {
            alert(error.message);
        });
    }
}

function clearErrorMessages() {
    const errorElements = document.querySelectorAll('.error-message');
    errorElements.forEach(element => {
        element.textContent = '';
    });
}
