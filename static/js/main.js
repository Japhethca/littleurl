function submitForm(event) {
    event.preventDefault();
    const form = new FormData(event.target)
    fetch("/", { method: 'POST', body: form })
        .then(res => res.text())
        .then(url => {
            showGeneratedLink(url);
            event.target.reset();
        })
        .catch(err => console.log(err));
}

function showGeneratedLink(url) {
    const linkWrapper = document.querySelector(".lu-link-wrapper");
    if (!url) {
        linkWrapper.setAttribute('style', 'visibility: hidden;')
        return
    }

    const linkTag = document.querySelector('.lu-gen-link');
    linkTag.href  = url;
    linkTag.textContent = url;
    linkWrapper.setAttribute('style', 'visibility: visible')
}