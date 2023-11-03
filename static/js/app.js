import Chat from "./components/Chat.js";
import Contact from "./components/Contact.js";
import CreatePostForm from "./components/CreatePostForm.js";
import LoginForm from "./components/LoginForm.js";
import Post from "./components/Post.js";
import SignupForm from "./components/SignupForm.js";

customElements.define("signup-form", SignupForm);
customElements.define("create-post-form", CreatePostForm);
customElements.define("login-form", LoginForm);
customElements.define("post-element", Post);
customElements.define("chat-element", Chat);
customElements.define("contact-element", Contact);

const codeListOnlineUsers = 1;
const codeUserLogin = 2;
const codeUserLogout = 3;
const codeNewPost = 4;
const codeNewComment = 5;
const codeNewDirectMessage = 6;

const nav = document.getElementById("nav");
const header = document.getElementById("header");
const main = document.getElementById("main");
const chat = document.getElementById("chat");

let user;
let socket;

const users = new Map();
const onlineUsers = new Map();
const posts = new Map();
const chats = new Map();

let contactsActive = true;
let activeChatID = 0;

window.addEventListener("load", async () => {
    if (!(await getMyUserAsync())) {
        renderLoginForm();
        return;
    }
    let usersData = await getUsersAsync();
    if (usersData) {
        for (const item of usersData) {
            users.set(item.id, new Contact(user, item));
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
                    const contact = new Contact(user, item);
                    onlineUsers.set(item.id, contact);
                }
                if (contactsActive) {
                    renderUsers();
                }
                break;
            case codeUserLogin:
                console.log("User Logged In:");
                console.log(body);
                const contact = new Contact(user, body.data);
                onlineUsers.set(body.data.id, contact);
                if (contactsActive) {
                    renderUsers();
                }
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
            case codeNewDirectMessage:
                console.log("New Direct Message:");
                console.log(body.data);
                let c;
                let chatId;
                if (body.data.senderID == user.id) {
                    c = chats.get(body.data.targetID);
                    if (!c) {
                        return;
                    }
                    chatId = body.data.targetID;
                } else {
                    c = chats.get(body.data.senderID);
                    if (!c) {
                        return;
                    }
                    chatId = body.data.senderID;
                }
                c.addMessageAsync(body.data);
                if (activeChatID == chatId) {
                    c.displayMessages();
                }
                break;
        }
    }
    socket.onerror = (e) => {
        console.log(e);
    }
}

export function renderUsers() {
    contactsActive = true;
    activeChatID = 0;
    chat.innerHTML = "";
    let onlineDiv = document.createElement("div");
    onlineDiv.classList.add("online-users");
    let onlineTitle = document.createElement("h3");
    onlineTitle.innerText = "Online Users";
    onlineDiv.appendChild(onlineTitle);
    let offlineDiv = document.createElement("div");
    offlineDiv.classList.add("offline-users");
    let offlineTitle = document.createElement("h3");
    offlineTitle.innerText = "Offline Users";
    offlineDiv.appendChild(offlineTitle);
    const keys = users.keys();
    for (const key of keys) {
        if (key == user.id) {
            continue;
        }
        const contact = users.get(key);
        if (onlineUsers.get(key)) {
            onlineDiv.appendChild(contact);
        } else {
            offlineDiv.appendChild(contact);
        }
    }
    chat.appendChild(onlineDiv);
    chat.appendChild(document.createElement("hr"));
    chat.appendChild(offlineDiv);
}

export function renderSignupForm() {
    header.innerHTML = "";
    main.innerHTML = "";
    main.appendChild(new SignupForm());
}

export function renderLoginForm() {
    header.innerHTML = "";
    main.innerHTML = "";
    main.appendChild(new LoginForm());
}

export function renderChat(sender, target) {
    let chatElem = chats.get(target.id);
    if (!chatElem) {
        chatElem = new Chat(sender, target, socket);
        chats.set(target.id, chatElem);
    }
    contactsActive = false;
    activeChatID = target.id;
    chat.innerHTML = "";
    
    chat.appendChild(chatElem);
}

export function buildTimeString(unixTime) {
    const cTime = new Date(unixTime*1000);
    const dateStr = `${cTime.getUTCDate()}/${cTime.getUTCMonth()+1}/${cTime.getUTCFullYear()}`;
    const min = cTime.getUTCMinutes();
    let minStr = `${min}`;
    if (min < 10) {
        minStr = `0${min}`;
    }
    // const sec = cTime.getUTCSeconds();
    // let secStr = `${sec}`;
    // if (sec < 10) {
    //     secStr = `0${sec}`;
    // }
    const timeStr = `${cTime.getUTCHours()}:${minStr} UTC`;
    return `${dateStr} at ${timeStr}`;
}