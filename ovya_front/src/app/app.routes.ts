import { Routes } from '@angular/router';
import { HomepageComponent } from './pages/homepage/homepage.component';
import { CreateVisiteComponent } from './pages/create-visite/create-visite.component';
import { VisiteListComponent } from './pages/visite-list/visite-list.component';

export const routes: Routes = [
    { path: 'homepage', component: HomepageComponent},
    { path: '', pathMatch: 'full', redirectTo: 'homepage'},
    { path: 'create_visite', component: CreateVisiteComponent },
    { path: 'visite', component: VisiteListComponent }
];
