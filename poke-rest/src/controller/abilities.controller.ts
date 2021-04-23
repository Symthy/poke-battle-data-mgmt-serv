import { Abilities, Ability } from '@/models/abilities';
import { AbilitiesService } from '@/provider/abilities.service';
import { Controller, DefaultValuePipe, Get, Param, Query } from '@nestjs/common';
import { ApiOkResponse, ApiQuery, ApiTags } from '@nestjs/swagger';

@ApiTags('Get Poke Base Data')
@Controller('abilities')
export class AbilitiesController {
  constructor(private readonly abilitiesService: AbilitiesService) { }

  @Get('')
  @ApiQuery({
    name: 'name',
    description: "ability name",
    required: false,
    type: String
  })
  @ApiOkResponse({ type: [Abilities] })
  async getAbilities(
    @Query('name', new DefaultValuePipe('')) name: string
  ): Promise<Abilities> {
    return this.abilitiesService.getAbilities();
  }

  @Get(':id')
  @ApiOkResponse({ type: [Ability] })
  async getAbilityById(@Param('id') id: string): Promise<Ability> {
    return this.abilitiesService.getAbilityById(id);
  }
}
