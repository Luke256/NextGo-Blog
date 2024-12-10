import Image from "next/image";

export default function Home() {

  const message = fetch("http://localhost:8080/hello", { cache: "no-store" })
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
    </div>
  );
}
