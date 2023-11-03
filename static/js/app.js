import LoginForm from "./components/LoginForm.js";
import Post from "./components/Post.js";
import SignupForm from "./components/SignupForm.js";
import { renderLoginForm, renderSignupForm } from "./render.js";

customElements.define("signup-form", SignupForm);
customElements.define("login-form", LoginForm);
customElements.define("post-element", Post);

const codeNewPost = 4;

let user;
let socket;
const posts = new Map();

window.addEventListener("load", async () => {
    let res = await fetch("/api/user");
    if (res.status !== 200) {
        renderLoginForm();
        return;
    }
    let data = await res.json();
    user = data;
    main.innerText = `Welcome ${data.nickname}`;
    let logout = document.createElement("a");
    logout.innerText = "Logout";
    logout.href="/logout";
    main.appendChild(logout);
    setupWebSocket();
    let postsData = await getPostsAsync();
    for (const item of postsData) {
        console.log(item);
        const post = new Post(item, socket);
        main.appendChild(post);
        posts.set(item.id, post);
    }
});

async function getPostsAsync() {
    const res = await fetch("/api/posts");
    if (res.status !== 200) {
        console.log(res.statusText);
        return null;
    }
    const data = await res.json();
    return data;
}

function setupWebSocket() {
    socket = new WebSocket("ws://localhost:8080/websocket");
    socket.onopen = () => {
        console.log("Socket opened.");
    }
    socket.onclose = () => {
        console.log("Socket closed.");
    }
    socket.onmessage = (e) => {
        // console.log(e.data);
        const message = JSON.parse(e.data);
        const body = JSON.parse(message.body);
        switch (body.code) {
            case codeNewPost:
                console.log(msg.body.body.id);
                console.log(msg.body.body.authorID);
                console.log(msg.body.body.content);
                console.log(msg.body.body.categories);
                console.log(msg.body.body.timestamp);
                break;
            case 5:
                console.log("New Comment - Data:");
                console.log(body.data);
                break;
        }
    }
    socket.onerror = (e) => {
        console.log(e);
    }
}