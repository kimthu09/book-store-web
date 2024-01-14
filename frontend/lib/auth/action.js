"use server"

import { auth, signIn, signOut } from "./auth";

export const login = async (props) => {
    const { email, password } = props;

    try {
        await signIn("credentials", { email, password });
    } catch (err) {
        console.log(err);

        if (err.message.includes("CredentialsSignin")) {
            return { error: "Invalid username or password" };
        }
        throw err;
    }
};

export const logOut = async () => {
    await signOut();
};

export const getApiKey = async () => {
    const session = await auth()
    return session?.user?.token
}

export const getUser = async () => {
    const session = await auth()
    return session?.user
}