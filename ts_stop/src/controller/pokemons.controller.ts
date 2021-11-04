import { Pokemon, Pokemons } from '@/models/pokemons';
import { PokemonsService } from '@/provider/pokemons.service';
import { Controller, DefaultValuePipe, Get, Param, Query } from '@nestjs/common';
import { ApiOkResponse, ApiQuery, ApiTags } from '@nestjs/swagger';

@ApiTags('Get Poke Base Data')
@Controller('pokemons')
export class PokemonsController {
  constructor(private readonly pokemonsService: PokemonsService) { }


  @Get('')
  @ApiQuery({
    name: 'name',
    description: "pokemon name",
    required: false,
    type: String
  })
  @ApiQuery({
    name: 'type',
    description: "pokemon type",
    required: false,
    type: String
  })
  @ApiOkResponse({ type: [Pokemons] })
  async getPokemons(
    @Query('name', new DefaultValuePipe('')) name: string,
    @Query('type', new DefaultValuePipe('')) type: string
  ): Promise<Pokemons> {
    return this.pokemonsService.getPokemons();
  }

  @Get(':id')
  @ApiOkResponse({ type: [Pokemon] })
  async getPokemonsById(@Param('id') id: string): Promise<Pokemon> {
    return this.pokemonsService.getPokemonById(id);
  }
}
