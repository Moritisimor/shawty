const aliasTextArea = document.getElementById("aliasTextArea");
const urlTextArea   = document.getElementById("urlTextArea");
const sendButton    = document.getElementById("sendButton");

const serverURL = window.location.href;

sendButton.addEventListener("click", () => {
    const alias = aliasTextArea.value;
    const url = urlTextArea.value;

    if (alias.trim() === "") {
        alert("Alias must not be empty or whitespace-only!");
        return;
    }

    if (url.trim() === "") {
        alert("URL must not be empty or whitespace-only!");
        return;
    }

    fetch(`${serverURL}api/alias`, {
        method: "POST",
        body: JSON.stringify({
            alias: alias,
            url: url
        })
    })
    .then((res) => {
        if (!res.ok) {
            if (res.status === 409) {
                throw new Error("An alias with this name already exists! Try another one");
            }

            throw new Error(`Error response from server: ${res.status}`);
        }

        return res.json();
    })
    .then((json) => alert(`Successfully posted alias.\nID: ${json["id"]}\nURL: ${serverURL}link/${alias}`))
    .catch((err) => alert(`Error while posting: ${err}`));
});
