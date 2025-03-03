document.addEventListener('htmx:afterRequest', focusOnSearch);

function focusOnSearch() {
    document.getElementById('search-input').focus()
}

function sort(newSortValue) {
    const sortElement = document.getElementById('sort');
    const orderElement = document.getElementById('order');
    const sortValue = sortElement.value;
    const orderValue = orderElement.value;

    let newOrderValue = orderValue

    if (newSortValue === sortValue) {
        newOrderValue = orderValue === 'desc' ? 'asc' : 'desc';
    }

    sortElement.value = newSortValue;
    orderElement.value = newOrderValue;

    document.getElementById('input-form').dispatchEvent(new CustomEvent('sort'));
}
