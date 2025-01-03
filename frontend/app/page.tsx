import Image from "next/image";
import Link from "next/link";

export default function Home() {
  const message = fetch(process.env.BACKEND_URL+"/hello", { cache: "no-store" })
    .then((response) => response.text())

  return (
    <div>
      <h1>Hello World</h1>
      <Image
        src="/vercel.svg"
        alt="Vercel Logo"
        width={72}
        height={16}
      />
      <p>{message}</p>
      <Link href="/login">Login</Link>
      <br />
      <Link href="/hoge">Hoge</Link>
    </div>
  );
}
