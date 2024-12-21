'use client';

import { useEffect } from "react";

const Login = () => {

    
    const genSession = async () => {
        const response = await fetch(process.env.NEXT_PUBLIC_BACKEND_URL+"/create-session", { cache: "no-store" });
        console.log(process.env.NEXT_PUBLIC_BACKEND_URL+"/create-session")
        const session = await response.json();
        console.log(session);
    }
    
    useEffect(() => {
        genSession();
    })

    return (
        <div>
        <h1>Login</h1>
        
        </div>
    );
}

export default Login;