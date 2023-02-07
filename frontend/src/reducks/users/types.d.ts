export type User = {
    id: number
    firstName: string
    lastName: string
    fullName: string
    emailAddress: string
}

export type SignUpField = {
    firstName: string
    lastName: string
    emailAddress: string
    password: string
    confirmPass: string
}

export type LogInField = {
    emailAddress: string
    password: string
}

export type ResponseUserData = {
    user: {
        id: number
        first_name: string
        last_name: string
        email: string
        full_name: string
        token: string
    }
}
