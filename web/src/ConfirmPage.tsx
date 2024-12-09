import React from "react";
import { useParams } from "react-router-dom";

function ConfirmPage() {
  const { token } = useParams();

  const handleConfirm = async () => {
    if (!token) {
      return;
    }

    const response = await fetch(
      `http://localhost:8080/v1/users/activate/${token}`,
      {
        method: "PUT",
      }
    );

    if (response.ok) {
      alert("Email confirmed successfully!");
    } else {
      alert("Failed to confirm email.");
    }
  };

  return (
    <div
      className="w-auto min-h-screen p-5 text-center justify-center items-center flex flex-col gap-4"
      style={{ backgroundColor: "#282c34", color: "white" }}
    >
      <h1 className=" text-5xl ">Confirm your email</h1>
      <p className="text-xl">
        Once you confirm your email, you will be able to log in to your account
        and start using GoSocial.
      </p>

      <button
        className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
        onClick={handleConfirm}
      >
        Confirm
      </button>
    </div>
  );
}

export default ConfirmPage;
