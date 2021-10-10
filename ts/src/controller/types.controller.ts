import { TypeCompatibility, TypeCompatibilityTable, Types } from '@/models/typeCompatibility';
import { Controller, Get, Param } from '@nestjs/common';
import { ApiOkResponse, ApiTags } from '@nestjs/swagger';
import { TypesService } from '../provider/types.service';

@ApiTags('Get Poke Base Data')
@Controller('types')
export class TypesController {
  constructor(private readonly typesService: TypesService) { }

  @Get('')
  @ApiOkResponse({ type: [Types] })
  async getTypes(): Promise<Types> {
    return this.typesService.getTypes();
  }

  @Get('compatibility')
  @ApiOkResponse({ type: [TypeCompatibilityTable] })
  async getTypeCompatibilityTable(): Promise<TypeCompatibilityTable> {
    return this.typesService.getTypeCompatibilityTable();
  }

  @Get('compatibility/attack/:type')
  @ApiOkResponse({ type: [TypeCompatibility] })
  async getTypeCompatibilityForAttack(@Param('type') type: string): Promise<TypeCompatibility> {
    return this.typesService.getTypeCompatibility(type);
  }

  @Get('compatibility/defence/:type')
  @ApiOkResponse({ type: [TypeCompatibility] })
  async getTypeCompatibilityForDefence(@Param('type') type: string): Promise<TypeCompatibility> {
    return this.typesService.getTypeCompatibility(type);
  }
}
