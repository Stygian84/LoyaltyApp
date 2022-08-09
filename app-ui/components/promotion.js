import Link from "next/link";
import React, { useEffect, useState } from "react";
import LabelContent from "../components/labelContent";
import axios from "axios";

export default function Promotion({ promo }) {
    axios.defaults.baseURL = "http://localhost:8080";

    // switch (promo.earn_rate_type) {
    //     case "add":
    //         //getProg(promo.program)
    //         var earnRate = 0;
    //         if (prog.initial_earn_rate != null) {
    //             earnRate = prog.initial_earn_rate;
    //             var offerInfo = "extra";
    //         }
    //         break;
    //     case "mul":
    //         //getProg(promo.program)
    //         var earnRate = 0;
    //         if (prog.initial_earn_rate != null) {
    //             earnRate = prog.initial_earn_rate * promo.constant;
    //             var offerInfo = "X";
    //         }
    //         break;
    // }
    //
    return (
        <div className="card rounded shadow flex flex-col items-center justify center p-5 border-4 border-yellow-500">
            <div className="flex flex-col items-center justify center p-5">
                <p className=" text-pink font-bold text-lg mb-4">
                    {promo.constant.toFixed(2)} points offered!
                </p>
                <LabelContent title="Program Name">
                    <h2 className="text-lg">{promo.name.String}</h2>
                </LabelContent>
                <LabelContent title="Promotion Type">
                    <p className="">{promo.promo_type} </p>
                </LabelContent>
                <LabelContent title="Point to Rewards Ratio">
                    <p className="">
                        1 point : {promo.initial_earn_rate.Float64.toFixed(2)}{" "}
                        {promo.currency_name.String}
                    </p>
                </LabelContent>
                <LabelContent title="Partner Code">
                    <p className="">{promo.partner_code.String}</p>
                </LabelContent>
                <LabelContent title="Estimated Transfer Time">
                    <p className="">Up to {promo.processing_time.String}</p>
                </LabelContent>

                <LabelContent title="End Date">
                    <p className="">{promo.end_date}</p>
                </LabelContent>
                <Link href={`/transaction/${promo.id_2.Int64}`}>
                    <button className="mt-4 text-lg cursor-pointer bg-blue-300 rounded drop-shadow-sm px-4 ">
                        <p>Transfer credits</p>
                    </button>
                </Link>
            </div>
        </div>
    );
}
