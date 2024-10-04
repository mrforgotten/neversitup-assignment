type getSomeType = () => string;

export const printSome = (input: string, getSome: getSomeType): string => {
    //do some manipulate
    return `${getSome()} with input ${input}`;
};
