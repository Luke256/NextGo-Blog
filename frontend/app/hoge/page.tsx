'use client';

import { useEffect, useState } from "react";

const HogePage: React.FC = () => {
    const [res, setRes] = useState("");

    const callBackend = async () => {
        const response = await fetch(process.env.NEXT_PUBLIC_APP_URL+"/sess/read-session", { cache: "no-store" , credentials: 'include' });
        const hoge = await response.text();
        setRes(hoge);
        if (response.status === 401) {
            setRes("Unauthorized");
        }
    }
    useEffect(() => {
        callBackend();
    })
    return <div>
        <h1>Hoge</h1>
        <p>{res}</p>
    </div>;
}

export default HogePage;