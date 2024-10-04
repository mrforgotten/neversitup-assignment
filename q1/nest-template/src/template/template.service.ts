import { Injectable } from '@nestjs/common';
import corefunc from 'src/core-func';

@Injectable()
export class TemplateService {
    constructor() {}

    getSome(): string {
        return 'some';
    }

    async getSomeInput(input: string): Promise<string> {
        return corefunc.printSome(input, this.getSome);
    }
}
