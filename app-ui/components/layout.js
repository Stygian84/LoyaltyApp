import Head from "next/head";
import Link from "next/link";
import React, { useContext, useEffect, useState } from "react";

import { useRouter } from "next/router";

export default function Layout({ title, children }) {
  const router = useRouter();

  return (
    <>
      <Head>
        <title>{title + "-ESC"}</title>
        <meta name="description" content="ecommerce website" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <div className="flex min-h-screen flex-col justify-between">
        <header>
          <nav className="flex h-12 items-centre px-4 py-2 justify-between shadow-md bg-blue-800">
            <Link href="/homepage">
              <a className="text-lg font-bold text-white">Transfer Connect </a>
            </Link>
          </nav>
        </header>

        <main className="my-auto  mt-0 ">{children}</main>
        <footer className="flex  h-10 justify-center items-center shadow-inner">
          Copyright © 2022{" "}
        </footer>
      </div>
    </>
  );
}