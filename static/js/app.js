import SignupForm from "./components/SignupForm.js";

customElements.define("signup-form", SignupForm);

const main = document.getElementById("main");

window.addEventListener("load", async () => {
    let res = await fetch("/api/user");
    if (res.status == 200) {
        let json = await res.json();
        main.innerText = `Welcome ${json.nickname}`;
    } else {
        main.appendChild(new SignupForm());
    }
})