// Global item counter
let itemCounter = 0;

// Initialize with one item on page load
document.addEventListener('DOMContentLoaded', function() {
    addItem();
});

// Add new item to the form
function addItem() {
    itemCounter++;
    const container = document.getElementById('itemsContainer');
    
    const itemCard = document.createElement('div');
    itemCard.className = 'item-card';
    itemCard.id = `item-${itemCounter}`;
    
    itemCard.innerHTML = `
        <div class="item-header">
            <h3>Item ${itemCounter}</h3>
            <button type="button" class="btn btn-danger" onclick="removeItem(${itemCounter})">
                Remove
            </button>
        </div>
        <div class="item-grid">
            <div class="form-group">
                <label for="itemName-${itemCounter}">Item Name *</label>
                <input type="text" id="itemName-${itemCounter}" placeholder="e.g., Laptop" required>
            </div>
            <div class="form-group">
                <label for="itemPrice-${itemCounter}">Price ($) *</label>
                <input type="number" id="itemPrice-${itemCounter}" placeholder="100.00" step="0.01" min="0" required>
            </div>
            <div class="form-group">
                <label for="itemQuantity-${itemCounter}">Quantity *</label>
                <input type="number" id="itemQuantity-${itemCounter}" placeholder="1" min="1" required>
            </div>
        </div>
        <div class="form-group full-width" style="margin-top: 12px;">
            <label for="itemDescription-${itemCounter}">Description (Optional)</label>
            <input type="text" id="itemDescription-${itemCounter}" placeholder="Product details">
        </div>
    `;
    
    container.appendChild(itemCard);
}

// Remove item from the form
function removeItem(itemId) {
    const itemCard = document.getElementById(`item-${itemId}`);
    if (itemCard) {
        itemCard.remove();
    }
    
    // Check if no items left, add one
    const container = document.getElementById('itemsContainer');
    if (container.children.length === 0) {
        addItem();
    }
}

// Reset the entire form
function resetForm() {
    // Reset address fields
    document.getElementById('state').value = '';
    document.getElementById('zipcode').value = '';
    document.getElementById('city').value = '';
    document.getElementById('street').value = '';
    
    // Clear all items
    const container = document.getElementById('itemsContainer');
    container.innerHTML = '';
    itemCounter = 0;
    
    // Add one item
    addItem();
    
    // Hide results and errors
    document.getElementById('resultsSection').style.display = 'none';
    document.getElementById('errorSection').style.display = 'none';
}

// Collect form data and validate
function getFormData() {
    const state = document.getElementById('state').value;
    const zipcode = document.getElementById('zipcode').value;
    const city = document.getElementById('city').value;
    const street = document.getElementById('street').value;
    
    // Validation
    if (!state) {
        throw new Error('Please select a state');
    }
    
    if (!zipcode) {
        throw new Error('Please enter a ZIP code');
    }
    
    // Collect items
    const items = [];
    const container = document.getElementById('itemsContainer');
    const itemCards = container.querySelectorAll('.item-card');
    
    if (itemCards.length === 0) {
        throw new Error('Please add at least one item');
    }
    
    itemCards.forEach((card, index) => {
        const id = card.id.split('-')[1];
        const name = document.getElementById(`itemName-${id}`)?.value;
        const price = parseFloat(document.getElementById(`itemPrice-${id}`)?.value);
        const quantity = parseInt(document.getElementById(`itemQuantity-${id}`)?.value);
        const description = document.getElementById(`itemDescription-${id}`)?.value;
        
        // Validation
        if (!name || !name.trim()) {
            throw new Error(`Item ${index + 1}: Please enter an item name`);
        }
        
        if (isNaN(price) || price < 0) {
            throw new Error(`Item ${index + 1}: Please enter a valid price (0 or greater)`);
        }
        
        if (isNaN(quantity) || quantity < 1) {
            throw new Error(`Item ${index + 1}: Please enter a valid quantity (1 or greater)`);
        }
        
        items.push({
            id: `item-${id}`,
            name: name.trim(),
            price: price,
            quantity: quantity,
            description: description?.trim() || ''
        });
    });
    
    return {
        address: {
            street: street.trim(),
            city: city.trim(),
            state: state.trim(),
            country: 'US',
            zipcode: zipcode.trim()
        },
        items: items
    };
}

// Calculate tax by calling the API
async function calculateTax() {
    // Hide previous results and errors
    document.getElementById('resultsSection').style.display = 'none';
    document.getElementById('errorSection').style.display = 'none';
    
    try {
        // Get and validate form data
        const formData = getFormData();
        
        // Show loading indicator
        document.getElementById('loadingIndicator').style.display = 'block';
        
        // Call API
        const response = await fetch('/api/v1/calculate-tax', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(formData)
        });
        
        // Hide loading indicator
        document.getElementById('loadingIndicator').style.display = 'none';
        
        if (!response.ok) {
            const errorData = await response.json();
            throw new Error(errorData.message || 'Failed to calculate tax');
        }
        
        const result = await response.json();
        displayResults(result);
        
    } catch (error) {
        // Hide loading indicator
        document.getElementById('loadingIndicator').style.display = 'none';
        
        // Show error
        showError(error.message);
    }
}

// Display calculation results
function displayResults(data) {
    // Update summary
    document.getElementById('subtotal').textContent = `$${data.subtotal.toFixed(2)}`;
    document.getElementById('totalTax').textContent = `$${data.total_tax.toFixed(2)}`;
    document.getElementById('grandTotal').textContent = `$${data.grand_total.toFixed(2)}`;
    document.getElementById('jurisdiction').textContent = data.tax_jurisdiction;
    
    // Display items breakdown
    const breakdownContainer = document.getElementById('itemsBreakdown');
    breakdownContainer.innerHTML = '';
    
    data.items.forEach((item, index) => {
        const breakdownItem = document.createElement('div');
        breakdownItem.className = 'breakdown-item';
        breakdownItem.innerHTML = `
            <div class="breakdown-item-header">
                <h4>${item.item_name}</h4>
                <strong>$${item.total_amount.toFixed(2)}</strong>
            </div>
            <div class="breakdown-details">
                <div>
                    <span class="label">Price:</span>
                    <span class="value">$${item.price.toFixed(2)}</span>
                </div>
                <div>
                    <span class="label">Quantity:</span>
                    <span class="value">${item.quantity}</span>
                </div>
                <div>
                    <span class="label">Subtotal:</span>
                    <span class="value">$${item.subtotal.toFixed(2)}</span>
                </div>
                <div>
                    <span class="label">Tax Rate:</span>
                    <span class="value">${item.tax_rate.toFixed(2)}%</span>
                </div>
                <div>
                    <span class="label">Tax Amount:</span>
                    <span class="value">$${item.tax_amount.toFixed(2)}</span>
                </div>
                <div>
                    <span class="label">Total:</span>
                    <span class="value">$${item.total_amount.toFixed(2)}</span>
                </div>
            </div>
        `;
        breakdownContainer.appendChild(breakdownItem);
    });
    
    // Show results section
    document.getElementById('resultsSection').style.display = 'block';
    
    // Scroll to results
    document.getElementById('resultsSection').scrollIntoView({ 
        behavior: 'smooth', 
        block: 'nearest' 
    });
}

// Show error message
function showError(message) {
    const errorSection = document.getElementById('errorSection');
    const errorMessage = document.getElementById('errorMessage');
    
    errorMessage.textContent = message;
    errorSection.style.display = 'block';
    
    // Scroll to error
    errorSection.scrollIntoView({ 
        behavior: 'smooth', 
        block: 'nearest' 
    });
}

// Handle enter key in form
document.addEventListener('keypress', function(e) {
    if (e.key === 'Enter' && e.target.tagName !== 'BUTTON') {
        e.preventDefault();
        calculateTax();
    }
});
