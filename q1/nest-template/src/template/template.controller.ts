import { Controller, Get, Param } from '@nestjs/common';
import { TemplateService } from './template.service';

@Controller('template')
export class TemplateController {
    constructor(private readonly service: TemplateService) {}

    @Get('/some')
    async getSome() {
        return this.service.getSome();
    }

    @Get('/some/:input')
    async getSomeInput(@Param('input') input: string) {
        return this.service.getSomeInput(input)
    }
}
