import { generateKeyPairSync } from "crypto";

export const getRSAKeyPair = () => {
    const { privateKey, publicKey } = generateKeyPairSync('rsa', {
        modulusLength: 2048,
    });
    console.log("done generating key pair1")
    const exportedPublicKeyBuffer = publicKey.export({
        type: "pkcs1",
        format: "pem",
    });
    // fs.writeFileSync("public.pem", exportedPublicKeyBuffer, { encoding: "utf-8" });

    const exportedPrivateKeyBuffer = privateKey.export({
        type: "pkcs1",
        format: "pem",
    });
    // fs.writeFileSync("private.pem", exportedPrivateKeyBuffer, {
    //     encoding: "utf-8",
    // });
    return { exportedPrivateKeyBuffer, exportedPublicKeyBuffer  };
}    