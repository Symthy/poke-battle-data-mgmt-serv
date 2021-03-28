import { ApiProperty } from "@nestjs/swagger";

export class Pokemon {

  @ApiProperty()
  pokedexId: string;

  @ApiProperty()
  name: string;

  @ApiProperty()
  englishName: string;

  @ApiProperty()
  baseStatsH: number;

  @ApiProperty()
  baseStatsA: number;

  @ApiProperty()
  baseStatsB: number;

  @ApiProperty()
  baseStatsC: number;

  @ApiProperty()
  baseStatsD: number;

  @ApiProperty()
  baseStatsS: number;

  @ApiProperty()
  type: Array<string>;
}

export class Pokemons {
  @ApiProperty({ type: [Pokemon] })
  pokemons: Array<Pokemon>;
}
