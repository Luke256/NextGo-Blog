'use client';

import { useEffect, useState } from "react";

const Login = () => {
    const [result, setResult] = useState("");
    
    const genSession = async () => {
        const response = await fetch(
            process.env.NEXT_PUBLIC_APP_URL+"/create-session",
            {
                cache: "no-store",
                credentials: "include",
            }
        );
        if (response.status === 200) {
            setResult("Session created");
            console.log(response);
        }
        // Error
        else {
            setResult("Error creating session");
        }
    }
    
    useEffect(() => {
        genSession();
    }, [])

    return (
        <div>
        <h1>Login</h1>
        <p>{result}</p>
        </div>
    );
}

export default Login;