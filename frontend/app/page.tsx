import { Shorten } from "./shortener";

export default function Home() {
  return (
    <>
      <div className="flex justify-center items-center p-10">
        <p className="text-3xl">
          URL Shortener
        </p>
      </div>
      <div className="flex justify-center h-full">
        <div>
          <Shorten></Shorten>
        </div>
      </div>
    </>
  );
}
