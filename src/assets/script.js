document.body.addEventListener('htmx:beforeRequest',  updateQuery);
document.addEventListener('DOMContentLoaded',  refreshContent);

window.addEventListener('popstate',  refreshContent);

document.getElementById("search").addEventListener("search-updated", () => {
    console.log("search-updated event triggered!");
});

var query = "";

function refreshContent() {
    console.log("refreshContent()")
    resolveQueryFromUrl();
    updateSearchInputValue();
}

function resolveQueryFromUrl() {
    console.log("resolveQueryFromUrl");
     query = window.location.search;
}

function updateQuery(event) {
    console.log("updateQuery");
    const windowSearch = new URLSearchParams(window.location.search);
    const requestParams = Object.entries(event.detail.requestConfig.parameters).map(e => {
        return { key: e[0], value: e[1] }
    });
    const searchParam = requestParams.find(x => x.key === 'search')

    if (searchParam?.value === ''){
        windowSearch.delete('search')
    }
    else
    {
        windowSearch.set(searchParam.key, searchParam.value)
    }

    let searchString = windowSearch.toString();
    query = (!searchString || searchString === '') ? '' : '?' + searchString;
    updateUrls()
}

function updateSearchInputValue() {
    console.log("updateSearchInputValue");
    const params = new URLSearchParams(query);
    const search = "search";
    let searchElement = document.getElementById(search)
    searchElement.value = params.get(search) ?? '';
    console.log("Firing reload event")
    searchElement.dispatchEvent(new CustomEvent("search-updated"));
}

function updateUrls() {
    console.log("updateUrls");
    var url = window.location.origin + query;
    window.history.pushState(null, null, url);
    updateRssLink(query);
}
    
function updateRssLink() {
    console.log("updateRssLink");
    const origin = window.location.origin;
    const rssElement = document.getElementById('rss-link');
    rssElement.setAttribute('href', origin + '/rss' + query)
}
