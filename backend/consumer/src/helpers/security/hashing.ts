import crypto from 'crypto';

export const generateHash = (message: string) => {
    const hash = crypto.createHash('sha256');
    hash.update(message);
    return hash.digest('hex');
}

export const generateRandomHash = () => {
    const randomMessage = crypto.randomBytes(64).toString('hex');
    return generateHash(randomMessage);
}

export const compareHash = (message: string, hash: string) => {
    return generateHash(message) === hash;
}