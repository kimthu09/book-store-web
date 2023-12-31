export const authConfig = {
    pages: {
        signIn: "/login",
    },
    providers: [],
    callbacks: {
        jwt({ token, user }) {
            // console.log(user)
            if (user) {
                token.id = user;
            }
            return token;
        },
        session({ session, token }) {
            if (token) {
                session.user = token.id;
            }
            return session;
        },
        async authorized({ auth, request }) {
            const user = auth?.user;
            // console.log(user)
            const isOnLoginPage = request.nextUrl?.pathname.startsWith("/login");

            if (user && isOnLoginPage) {
                return Response.redirect(new URL("/", request.nextUrl));
            }
            if (!user && !isOnLoginPage) {
                return false
            }
            return true
        },
    },
};
