import Link from "next/link";
import React from "react";
import Image from "next/dist/client/image";
import LabelContent from "../components/labelContent";

const Program = ({ program }) => {
    return (
        <div className="card rounded shadow flex flex-col items-center justify center p-5  border">
            {/* <Image 
            src= {program.image}
            alt={program.description}
            className="rounded shadow items-center "
            width="100%"
            height="100%"
            >
            </Image> */}

            <div className="flex flex-col items-center justify center p-5">
                <LabelContent title="Program Name">
                    <h2 className="text-lg">{program.name}</h2>
                </LabelContent>
                <LabelContent title="Point to Rewards Ratio">
                    <p className="">
                        1 point : {program.initial_earn_rate.toFixed(2)}{" "}
                        {program.currency_name}
                    </p>
                </LabelContent>
                <LabelContent title="Partner Code">
                    <p className="">{program.partner_code}</p>
                </LabelContent>
                <LabelContent title="Estimated Transfer Time">
                    <p className="">Up to {program.processing_time}</p>
                </LabelContent>
                <Link href={`/transaction/${program.id}`}>
                    <button className="mt-4 text-lg cursor-pointer bg-blue-300 rounded drop-shadow-sm px-4 ">
                        <p>Transfer credits</p>
                    </button>
                </Link>
            </div>
        </div>
    );
};
export default Program;
