"use server"

import { signIn, signOut } from "./auth";

export const login = async (prevState, formData) => {
    const { email, password } = Object.fromEntries(formData);

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