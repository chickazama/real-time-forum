import CreatePostForm from "./components/CreatePostForm.js";
import LoginForm from "./components/LoginForm.js";
import Post from "./components/Post.js";
import SignupForm from "./components/SignupForm.js";
import { renderLoginForm, renderSignupForm } from "./render.js";

customElements.define("signup-form", SignupForm);
customElements.define("create-post-form", CreatePostForm);
customElements.define("login-form", LoginForm);
customElements.define("post-element", Post);

const codeListOnlineUsers = 1;
const codeUserLogin = 2;
const codeUserLogout = 3;
const codeNewPost = 4;
const codeNewComment = 5;

const header = document.getElementById("header");
const main = document.getElementById("main");
const chat = document.getElementById("chat");

let user;
let socket;

const users = new Map();
const onlineUsers = new Map();
const posts = new Map();

window.addEventListener("load", async () => {
    if (!(await getMyUserAsync())) {
        renderLoginForm();
        return;
    }
    let usersData = await getUsersAsync();
    if (usersData) {
        for (const item of usersData) {
            users.set(item.id, item);
        }
    }
    let logout = document.createElement("a");
    logout.innerText = "Logout";
    logout.href="/logout";
    header.appendChild(logout);
    setupWebSocket();
    let postForm = new CreatePostForm(user, socket);
    header.appendChild(postForm);
    let postsData = await getPostsAsync();
    for (const item of postsData) {
        console.log(item);
        const post = new Post(user, item, socket);
        main.appendChild(post);
        posts.set(item.id, post);
    }
});

async function getUsersAsync() {
    let res = await fetch("/api/users");
    if (res.status !== 200) {
        return null;
    }
    let data = await res.json();
    if (!data) {
        return null;
    }
    return data;
}

async function getMyUserAsync() {
    let res = await fetch("/api/user");
    if (res.status !== 200) {
        return false;
    }
    let data = await res.json();
    if (!data) {
        return false;
    }
    user = data;
    return true;
}

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
    socket.onmessage = async (e) => {
        // console.log(e.data);
        const message = JSON.parse(e.data);
        const body = JSON.parse(message.body);
        switch (body.code) {
            case codeListOnlineUsers:
                console.log("Online Users:");
                console.log(body.data);
                for (const item of body.data) {
                    onlineUsers.set(item.id, item);
                }
                renderUsers();
                break;
            case codeUserLogin:
                console.log("User Logged In:");
                console.log(body);
                onlineUsers.set(body.data.id, body.data);
                renderUsers();
                break;
            case codeUserLogout:
                console.log("User Logged out:");
                console.log(body);
                onlineUsers.set(body.data.id, null);
                renderUsers();
                break;
            case codeNewPost:
                console.log("New Post:");
                console.log(body.data);
                const newPost = new Post(user, body.data, socket);
                posts.set(body.data.id, newPost);
                main.prepend(newPost);
                break;
            case codeNewComment:
                console.log("New Comment:");
                console.log(body.data);
                const post = posts.get(body.data.postID);
                await post.addCommentAsync(body.data);
                if (post.getAttribute("active") == "true") {
                    post.setAttribute("active", "refresh");
                }
                break;
        }
    }
    socket.onerror = (e) => {
        console.log(e);
    }
}

function renderUsers() {
    chat.innerHTML = "";
    let onlineDiv = document.createElement("div");
    let offlineDiv = document.createElement("div");
    const keys = users.keys();
    for (const key of keys) {
        const u = users.get(key);
        const p = document.createElement("p");
        p.innerText = u.nickname;
        if (onlineUsers.get(key)) {
            onlineDiv.appendChild(p);
        } else {
            offlineDiv.appendChild(p);
        }
    }
    chat.appendChild(onlineDiv);
    chat.appendChild(document.createElement("hr"));
    chat.appendChild(offlineDiv);
}