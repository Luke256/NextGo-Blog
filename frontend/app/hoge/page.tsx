'use client';

import { useEffect } from "react";

const HogePage: React.FC = () => {
    const callBackend = async () => {
        const response = await fetch(process.env.NEXT_PUBLIC_BACKEND_URL+"/sess/read-session", { cache: "no-store" , credentials: 'include' });
        const hoge = await response.json();
        console.log(hoge);
    }
    useEffect(() => {
        callBackend();
    })
    return <div>
        <h1>Hoge</h1>
    </div>;
}

export default HogePage;