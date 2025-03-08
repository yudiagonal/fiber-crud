document.addEventListener('DOMContentLoaded', () => {
    loadProducts();

    const form = document.getElementById('productForm');
    form.addEventListener('submit', (e) => {
        e.preventDefault();
        saveProduct();
    });
});

// Fungsi untuk memuat daftar produk
async function loadProducts() {
    try {
        const response = await fetch('/products');
        if (!response.ok) {
            throw new Error('Failed to load products');
        }
        const products = await response.json();
        console.log('Products from API:', products); // Debugging

        const tbody = document.querySelector('#productTable tbody');
        tbody.innerHTML = '';

        products.forEach(product => {
            const row = document.createElement('tr');
            row.innerHTML = `
                <td>${product.id}</td>
                <td>${product.name}</td>
                <td>${product.price}</td>
                <td class="actions">
                    <button class="edit" onclick="editProduct(${product.id})">Edit</button>
                    <button class="delete" onclick="deleteProduct(${product.id})">Delete</button>
                </td>
            `;
            tbody.appendChild(row);
        });
    } catch (error) {
        alert(error.message);
    }
}

// Fungsi untuk menyimpan produk (tambah/edit)
async function saveProduct() {
    const form = document.getElementById('productForm');
    const product = {
        id: form.productId.value || null,
        name: form.name.value,
        price: parseFloat(form.price.value)
    };

    const url = product.id ? `/products/${product.id}` : '/products';
    const method = product.id ? 'PUT' : 'POST';

    try {
        const response = await fetch(url, {
            method: method,
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(product)
        });

        if (!response.ok) {
            throw new Error('Failed to save product');
        }

        form.reset();
        loadProducts(); // Memuat ulang daftar produk setelah berhasil menyimpan
    } catch (error) {
        alert(error.message);
    }
}

// Fungsi untuk mengisi form dengan data produk yang akan diedit
async function editProduct(id) {
    try {
        const response = await fetch(`/products/${id}`);
        if (!response.ok) {
            throw new Error('Product not found');
        }
        const product = await response.json();

        document.getElementById('productId').value = product.id;
        document.getElementById('name').value = product.name;
        document.getElementById('price').value = product.price;
    } catch (error) {
        alert(error.message);
    }
}

// Fungsi untuk menghapus produk
async function deleteProduct(id) {
    try {
        const response = await fetch(`/products/${id}`, {
            method: 'DELETE'
        });
        if (!response.ok) {
            throw new Error('Failed to delete product');
        }
        loadProducts(); // Memuat ulang daftar produk setelah berhasil menghapus
    } catch (error) {
        alert(error.message);
    }
}