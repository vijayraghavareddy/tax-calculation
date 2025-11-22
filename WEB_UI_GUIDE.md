# Web UI Guide

## Overview

The Tax Calculation API includes a beautiful, modern web interface that makes calculating taxes simple and intuitive. No coding or API knowledge required!

## Getting Started

1. **Start the Server**
   ```bash
   go run main.go
   ```

2. **Open Your Browser**
   Navigate to: `http://localhost:8080`

3. **Start Calculating!**

## Interface Walkthrough

### 1. Header Section
- **Title:** Tax Calculator with emoji 🧾
- **Subtitle:** Explains the purpose
- **Gradient Background:** Purple gradient for modern look

### 2. Address Information Card

Fill in your address details:

**Required Fields:**
- ✅ **Country** - Dropdown with flag emojis (🇺🇸 🇨🇦 🇬🇧 🇩🇪 🇫🇷 🇮🇳 🇦🇺 🇯🇵)
- ✅ **ZIP/Postal Code** - Enter your area code

**Optional Fields:**
- State/Province
- City
- Street Address

> **Tip:** Only Country and ZIP code are required to calculate tax!

### 3. Items Section

Add products or services to calculate tax on:

**For Each Item:**
- **Item Name** - What you're buying (Required)
- **Price** - Unit price in dollars (Required)
- **Quantity** - How many units (Required)
- **Description** - Optional product details

**Actions:**
- Click **"+ Add Item"** to add more items
- Click **"Remove"** to delete an item
- At least one item is always required

### 4. Action Buttons

- **Calculate Tax** (Blue) - Submit and calculate
- **Reset** (White) - Clear all fields and start over

### 5. Results Section

After clicking "Calculate Tax", you'll see:

#### Summary Box (Green Background)
- **Subtotal:** Total before tax
- **Total Tax:** Tax amount calculated
- **Grand Total:** Final amount to pay (large, bold)
- **Tax Jurisdiction:** Location used for calculation

#### Items Breakdown
Detailed view for each item showing:
- Item name and total
- Original price
- Quantity purchased
- Subtotal (price × quantity)
- Tax rate applied (as percentage)
- Tax amount
- Item total (including tax)

## Example Usage

### Scenario: Buying Electronics in New York

1. **Address Section:**
   - Country: 🇺🇸 United States
   - ZIP Code: 10001
   - State: NY
   - City: New York

2. **Items:**
   
   **Item 1:**
   - Name: MacBook Pro
   - Price: 2499.00
   - Quantity: 1
   
   **Item 2:**
   - Name: Magic Mouse
   - Price: 79.00
   - Quantity: 2

3. **Click "Calculate Tax"**

4. **Results Show:**
   - Subtotal: $2,657.00
   - Total Tax: $226.23 (example, varies)
   - Grand Total: $2,883.23
   - Tax Jurisdiction: NY, US

## Features Highlight

### 🎨 Beautiful Design
- Modern gradient backgrounds
- Card-based layout
- Smooth animations
- Clean typography

### ✅ Real-time Validation
- Required fields highlighted
- Helpful error messages
- Prevents invalid submissions

### 📱 Mobile Responsive
- Works on phones and tablets
- Touch-friendly buttons
- Optimized layouts for small screens

### ⚡ Fast & Intuitive
- Instant calculations
- Loading indicators
- Smooth scrolling to results
- No page refreshes

### 🌍 Multi-Country Support
Country dropdown includes:
- 🇺🇸 United States (0-12% sales tax)
- 🇨🇦 Canada (5-15% GST/HST)
- 🇬🇧 United Kingdom (20% VAT)
- 🇩🇪 Germany (19% VAT)
- 🇫🇷 France (20% VAT)
- 🇮🇳 India (5-28% GST)
- 🇦🇺 Australia (10% GST)
- 🇯🇵 Japan (10% consumption tax)

## Error Handling

The UI provides clear, helpful error messages:

### Common Errors:
- ❌ "Please select a country"
- ❌ "Please enter a ZIP/Postal code"
- ❌ "Item 1: Please enter an item name"
- ❌ "Item 2: Please enter a valid price"
- ❌ "Item 1: Please enter a valid quantity"

Error messages appear in a red alert box above the results section.

## Keyboard Shortcuts

- **Enter** - Press anywhere in the form to submit
- **Tab** - Navigate between fields
- **Esc** - Close focused field (browser default)

## Tips for Best Experience

1. **Start Simple** - Enter just country and ZIP code first
2. **Add Items Gradually** - Start with one item, add more as needed
3. **Use Descriptions** - Optional but helpful for tracking
4. **Check Breakdown** - Review per-item tax to understand costs
5. **Save Results** - Screenshot or copy values before resetting

## Troubleshooting

### Page Won't Load
- Check server is running: `go run main.go`
- Verify URL: `http://localhost:8080`
- Check firewall settings

### Form Won't Submit
- Ensure all required fields (*) are filled
- Check price is positive number
- Verify quantity is at least 1
- Look for red error messages

### Results Not Showing
- Check browser console (F12) for errors
- Verify API is responding: `http://localhost:8080/api/v1/health`
- Try different browser

### Styling Looks Broken
- Clear browser cache (Ctrl+Shift+R or Cmd+Shift+R)
- Check static files exist in `/static/` folder
- Verify CSS file loaded in browser DevTools

## Customization

### Change Colors
Edit `static/styles.css` and modify CSS variables:
```css
:root {
    --primary-color: #4f46e5;  /* Blue */
    --secondary-color: #10b981; /* Green */
    /* ... more colors */
}
```

### Add More Countries
Edit `static/index.html` and add to the country dropdown:
```html
<option value="ES">🇪🇸 Spain</option>
```

Then update `services/tax_service.go` to handle the new country.

### Modify Layout
Edit `static/styles.css` for styling changes or `static/index.html` for structure changes.

## Browser Support

Tested and working on:
- ✅ Chrome/Edge (latest)
- ✅ Firefox (latest)
- ✅ Safari (latest)
- ✅ Mobile browsers

## Accessibility

The UI includes:
- Semantic HTML elements
- Label associations for form fields
- Keyboard navigation support
- Focus indicators
- Color contrast compliance

## API Integration

The UI communicates with the backend API at:
```
POST /api/v1/calculate-tax
```

JavaScript (`app.js`) handles:
- Form validation
- API requests (fetch)
- Error handling
- Results display
- Dynamic item management

## Files Structure

```
static/
├── index.html    # Main HTML structure
├── styles.css    # All CSS styling
└── app.js        # JavaScript logic
```

All files are served by the Go backend at startup.

---

## Need Help?

- **Documentation:** See README.md and API_DOCUMENTATION.md
- **API Testing:** Use Postman collection
- **Issues:** Check browser console for errors
- **Questions:** Open GitHub issue

Enjoy calculating taxes with style! 🎉
