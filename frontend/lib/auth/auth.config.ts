export const authConfig = {
    pages: {
        signIn: "/login",
    },
    providers: [],
    callbacks: {
        jwt({ token, user }) {
            // console.log(user)
            if (user) {
                token.id = user.data.id;
            }
            return token;
        },
        session({ session, token }) {
            if (token) {
                session.user.id = token.id;
            }
            return session;
        },
        async authorized({ auth, request }) {
            console.log(auth)
            const user = auth?.user;
            const isOnLoginPage = request.nextUrl?.pathname.startsWith("/login");

            if (user && isOnLoginPage) {
                return Response.redirect(new URL("/", request.nextUrl));
            }
            if (!user) {
                return false
            }
            return true
        },
    },
};
