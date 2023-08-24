<?php

namespace Database\Seeders;

use Illuminate\Database\Console\Seeds\WithoutModelEvents;
use Illuminate\Database\Seeder;

class Keys extends Seeder
{
    /**
     * Run the database seeds.
     */
    public function run()
    {
        \App\Models\Keys::factory(10)->create();

        \App\Models\Keys::factory()->create([
            'name' => 'Test Keys',
            'email' => 'test@example.com',
        ]);
    }
}
