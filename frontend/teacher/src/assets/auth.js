const login = document.getElementById("login")
const password = document.getElementById("password")

// registration page
const fName = document.getElementById("first-name")
const lName = document.getElementById("last-name")
const email = document.getElementById("email")
const regToken = document.getElementById("registration-token")
const confirmPassword = document.getElementById("confirm-password")

const errorMessage = document.getElementById("error-message")

// main page
const signout = document.getElementById("signout")


let submit = document.getElementById("submit") || document.getElementById("reg-submit")

if (submit !== null) {
    submit.addEventListener("click", async () => {
        let body = ''
        let url = ''
        if (submit.id === "submit") {
            url = "/auth/signin"
            body = JSON.stringify({
                login: login.value,
                password: password.value
            })
        } else {
            url = "/auth/registration"
            body = JSON.stringify({
                login: login.value,
                email: email.value,
                firstName: fName.value,
                lastName: lName.value,
                regToken: regToken.value,
                password: password.value
            })
        }

        errorMessage.style.display = 'none'
        const response = await fetch(url, {
            method: "POST",
            body
        })

        const json = await response.json()

        if (json.error !== undefined && json.error !== '') {
            errorMessage.style.display = 'block'
            errorMessage.innerText = json.error
        }

        if (json.success) {
            window.location.href = json.isTeacher ? "/teacher" : "/student"
        }
    })
}

if (signout !== null) {
    signout.addEventListener("click", async () => {
        const response = await fetch("/auth/signout", {
            method: "GET"
        })

        const json = await response.json()

        if (json.success) {
            window.location.href = "/"
        }
    })
}